from __future__ import annotations

from dataclasses import dataclass
from enum import Enum
from typing import Any, Dict, Iterable, List, Optional, Protocol, Sequence


class DBStore(Protocol):
    def insert(self, table: str, primary_key: str, values: Dict[str, Any]) -> Dict[str, Any]:
        ...

    def update(self, table: str, primary_key: str, values: Dict[str, Any]) -> None:
        ...

    def delete(self, table: str, primary_key: str, values: Dict[str, Any]) -> None:
        ...

    def select(
        self,
        table: str,
        filters: Sequence["Filter"],
        projection: Sequence[str],
        orderings: Sequence["OrderClause"],
        limit: Optional[int],
        offset: Optional[int],
    ) -> List[Dict[str, Any]]:
        ...


class Operator(str, Enum):
    EQ = "="
    NEQ = "!="
    IN = "IN"


class OrderDirection(str, Enum):
    ASC = "ASC"
    DESC = "DESC"


@dataclass(frozen=True)
class Filter:
    field: str
    operator: Operator
    value: Any


@dataclass(frozen=True)
class OrderClause:
    column: str
    direction: OrderDirection


_default_store: Optional[DBStore] = None


def set_default_store(store: DBStore) -> None:
    global _default_store
    _default_store = store


def _resolve_store(store_override: Optional[DBStore]) -> DBStore:
    if store_override is not None:
        return store_override
    if _default_store is None:
        raise RuntimeError("pdlinfra: no default DBStore configured")
    return _default_store


class Row:
    def __init__(
        self,
        table: str,
        primary_key: str,
        column_map: Dict[str, str],
        store: Optional[DBStore] = None,
    ) -> None:
        self._table = table
        self._primary_key = primary_key
        self._column_map = dict(column_map)
        self._store = store

    def set_store(self, store: DBStore) -> None:
        self._store = store

    @property
    def store(self) -> Optional[DBStore]:
        return self._store

    @property
    def table_name(self) -> str:
        return self._table

    @property
    def primary_key_name(self) -> str:
        return self._primary_key

    def column_map(self) -> Dict[str, str]:
        return dict(self._column_map)

    def values(self) -> Dict[str, Any]:
        result: Dict[str, Any] = {}
        for attr, column in self._column_map.items():
            result[column] = getattr(self, attr, None)
        return result


def hydrate(row: Row, values: Dict[str, Any]) -> None:
    if values is None:
        return
    for attr, column in row.column_map().items():
        if column in values:
            setattr(row, attr, values[column])
        elif attr in values:
            setattr(row, attr, values[attr])


class RowExecutor:
    @staticmethod
    def _resolve_store(row: Row) -> DBStore:
        store = _resolve_store(row.store)
        row.set_store(store)
        return store

    @staticmethod
    def create(row: Row) -> None:
        store = RowExecutor._resolve_store(row)
        values = row.values()
        inserted = store.insert(row.table_name, row.primary_key_name, values)
        if inserted:
            hydrate(row, inserted)

    @staticmethod
    def update(row: Row) -> None:
        store = RowExecutor._resolve_store(row)
        store.update(row.table_name, row.primary_key_name, row.values())

    @staticmethod
    def delete(row: Row) -> None:
        store = RowExecutor._resolve_store(row)
        store.delete(row.table_name, row.primary_key_name, row.values())

    @staticmethod
    def multi_insert_rows(*rows: Row) -> None:
        if not rows:
            return
        first = rows[0]
        store = RowExecutor._resolve_store(first)
        table = first.table_name
        primary_key = first.primary_key_name
        for row in rows:
            if row.table_name != table:
                raise ValueError("pdlinfra: mixed tables not supported in multi_insert_rows")
            inserted = store.insert(table, primary_key, row.values())
            if inserted:
                hydrate(row, inserted)


class QueryBuilder:
    def __init__(self, table: str, store: Optional[DBStore] = None) -> None:
        self._table = table
        self._store = store
        self._filters: List[Filter] = []
        self._projection: List[str] = []
        self._orderings: List[OrderClause] = []
        self._limit: Optional[int] = None
        self._offset: Optional[int] = None

    def with_store(self, store: DBStore) -> "QueryBuilder":
        self._store = store
        return self

    def eq(self, field: str, value: Any) -> "QueryBuilder":
        return self.filter(field, Operator.EQ, value)

    def neq(self, field: str, value: Any) -> "QueryBuilder":
        return self.filter(field, Operator.NEQ, value)

    def in_(self, field: str, values: Iterable[Any]) -> "QueryBuilder":
        return self.filter(field, Operator.IN, values)

    def filter(self, field: str, operator: Operator, value: Any) -> "QueryBuilder":
        self._filters.append(Filter(field, operator, value))
        return self

    def project(self, *columns: str) -> "QueryBuilder":
        self._projection.extend(columns)
        return self

    def order_by(self, column: str, direction: OrderDirection) -> "QueryBuilder":
        if not column:
            return self
        self._orderings.append(OrderClause(column, direction))
        return self

    def asc(self, column: str) -> "QueryBuilder":
        return self.order_by(column, OrderDirection.ASC)

    def desc(self, column: str) -> "QueryBuilder":
        return self.order_by(column, OrderDirection.DESC)

    def limit(self, value: int) -> "QueryBuilder":
        self._limit = value
        return self

    def offset(self, value: int) -> "QueryBuilder":
        self._offset = value
        return self

    def range(self, offset_value: int, limit_value: int) -> "QueryBuilder":
        self._offset = offset_value
        self._limit = limit_value
        return self

    def load(self) -> List[Dict[str, Any]]:
        store = _resolve_store(self._store)
        return store.select(
            self._table,
            list(self._filters),
            list(self._projection),
            list(self._orderings),
            self._limit,
            self._offset,
        )

    def delete(self, primary_key: str) -> None:
        store = _resolve_store(self._store)
        rows = store.select(self._table, list(self._filters), [primary_key], [], None, None)
        for entry in rows:
            store.delete(self._table, primary_key, entry)


__all__ = [
    "DBStore",
    "Filter",
    "Operator",
    "OrderClause",
    "OrderDirection",
    "QueryBuilder",
    "Row",
    "RowExecutor",
    "hydrate",
    "set_default_store",
]
