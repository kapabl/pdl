use std::collections::HashMap;
use std::fmt::{self, Display};
use std::sync::{Arc, RwLock};

use serde::{Deserialize, Serialize};
use serde_json::Value;

static DEFAULT_STORE: RwLock<Option<Arc<dyn DBStore>>> = RwLock::new(None);

pub type RowMap = HashMap<String, Value>;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct DbError {
    message: String,
}

impl DbError {
    pub fn new(message: impl Into<String>) -> Self {
        Self { message: message.into() }
    }
}

impl Display for DbError {
    fn fmt(&self, formatter: &mut fmt::Formatter<'_>) -> fmt::Result {
        formatter.write_str(&self.message)
    }
}

impl std::error::Error for DbError {}

pub trait DBStore: Send + Sync {
    fn insert(&self, table: &str, primary_key: &str, values: &mut RowMap) -> Result<RowMap, DbError>;
    fn update(&self, table: &str, primary_key: &str, values: &RowMap) -> Result<(), DbError>;
    fn delete(&self, table: &str, primary_key: &str, values: &RowMap) -> Result<(), DbError>;
    fn select(
        &self,
        table: &str,
        filters: &[Filter],
        projection: &[String],
        orderings: &[OrderClause],
        limit: Option<i32>,
        offset: Option<i32>,
    ) -> Result<Vec<RowMap>, DbError>;
}

#[derive(Debug, Clone)]
pub struct Filter {
    pub field: String,
    pub op: Operator,
    pub value: Value,
}

#[derive(Debug, Clone, Copy)]
pub enum Operator {
    Eq,
    Neq,
    In,
}

#[derive(Debug, Clone, Copy)]
pub enum OrderDirection {
    Asc,
    Desc,
}

#[derive(Debug, Clone)]
pub struct OrderClause {
    pub column: String,
    pub direction: OrderDirection,
}

pub fn set_default_store(store: Arc<dyn DBStore>) {
    let mut guard = DEFAULT_STORE.write().expect("default store poisoned");
    *guard = Some(store);
}

fn resolve_store(override_store: Option<Arc<dyn DBStore>>) -> Result<Arc<dyn DBStore>, DbError> {
    if let Some(store) = override_store {
        return Ok(store);
    }
    let guard = DEFAULT_STORE.read().expect("default store poisoned");
    guard.clone().ok_or_else(|| DbError::new("pdl_infrastructure: no DB store configured"))
}

#[derive(Debug, Clone)]
pub struct Row {
    table: String,
    primary_key: String,
    store: Option<Arc<dyn DBStore>>,
}

impl Row {
    pub fn new(table: impl Into<String>, primary_key: impl Into<String>) -> Self {
        Self {
            table: table.into(),
            primary_key: primary_key.into(),
            store: None,
        }
    }

    pub fn table(&self) -> &str {
        &self.table
    }

    pub fn primary_key(&self) -> &str {
        &self.primary_key
    }

    pub fn store(&self) -> Option<Arc<dyn DBStore>> {
        self.store.clone()
    }

    pub fn set_store(&mut self, store: Arc<dyn DBStore>) {
        self.store = Some(store);
    }
}

pub trait Record {
    fn row_meta(&self) -> &Row;
    fn row_meta_mut(&mut self) -> &mut Row;
    fn collect_values(&self) -> RowMap;
    fn apply_row_map(&mut self, values: &RowMap) -> Result<(), DbError>;
}

#[derive(Clone)]
pub struct QueryBuilder {
    table: String,
    store: Option<Arc<dyn DBStore>>,
    filters: Vec<Filter>,
    projection: Vec<String>,
    orderings: Vec<OrderClause>,
    limit: Option<i32>,
    offset: Option<i32>,
}

impl QueryBuilder {
    pub fn new(table: impl Into<String>, store: Option<Arc<dyn DBStore>>) -> Self {
        Self {
            table: table.into(),
            store,
            filters: Vec::new(),
            projection: Vec::new(),
            orderings: Vec::new(),
            limit: None,
            offset: None,
        }
    }

    pub fn with_store(&mut self, store: Arc<dyn DBStore>) -> &mut Self {
        self.store = Some(store);
        self
    }

    pub fn filter<T>(&mut self, field: impl Into<String>, op: Operator, value: T) -> &mut Self
    where
        T: Serialize,
    {
        let encoded = serde_json::to_value(value).unwrap_or(Value::Null);
        self.filters.push(Filter {
            field: field.into(),
            op,
            value: encoded,
        });
        self
    }

    pub fn eq<T>(&mut self, field: impl Into<String>, value: T) -> &mut Self
    where
        T: Serialize,
    {
        self.filter(field, Operator::Eq, value)
    }

    pub fn projection(&mut self, columns: &[&str]) -> &mut Self {
        self.projection
            .extend(columns.iter().map(|entry| entry.to_string()));
        self
    }

    pub fn order_by(&mut self, column: impl Into<String>, direction: OrderDirection) -> &mut Self {
        self.orderings.push(OrderClause {
            column: column.into(),
            direction,
        });
        self
    }

    pub fn limit(&mut self, value: i32) -> &mut Self {
        self.limit = Some(value);
        self
    }

    pub fn offset(&mut self, value: i32) -> &mut Self {
        self.offset = Some(value);
        self
    }

    pub fn range(&mut self, offset: i32, limit: i32) -> &mut Self {
        self.offset = Some(offset);
        self.limit = Some(limit);
        self
    }

    pub fn load(&self) -> Result<Vec<RowMap>, DbError> {
        let store = resolve_store(self.store.clone())?;
        store.select(
            &self.table,
            &self.filters,
            &self.projection,
            &self.orderings,
            self.limit,
            self.offset,
        )
    }

    pub fn delete(&self, primary_key: &str) -> Result<(), DbError> {
        let store = resolve_store(self.store.clone())?;
        let rows = store.select(
            &self.table,
            &self.filters,
            &[primary_key.to_string()],
            &[],
            None,
            None,
        )?;
        for row in rows {
            store.delete(&self.table, primary_key, &row)?;
        }
        Ok(())
    }
}

pub struct RowExecutor;

impl RowExecutor {
    pub fn create<T>(record: &mut T) -> Result<(), DbError>
    where
        T: Record,
    {
        let mut values = record.collect_values();
        let store = resolve_store(record.row_meta().store.clone())?;
        let inserted = store.insert(record.row_meta().table(), record.row_meta().primary_key(), &mut values)?;
        record.row_meta_mut().set_store(store);
        record.apply_row_map(&values)?;
        if !inserted.is_empty() {
            record.apply_row_map(&inserted)?;
        }
        Ok(())
    }

    pub fn update<T>(record: &mut T) -> Result<(), DbError>
    where
        T: Record,
    {
        let map = record.collect_values();
        let store = resolve_store(record.row_meta().store.clone())?;
        record.row_meta_mut().set_store(store.clone());
        store.update(record.row_meta().table(), record.row_meta().primary_key(), &map)
    }

    pub fn delete<T>(record: &mut T) -> Result<(), DbError>
    where
        T: Record,
    {
        let map = record.collect_values();
        let store = resolve_store(record.row_meta().store.clone())?;
        record.row_meta_mut().set_store(store.clone());
        store.delete(record.row_meta().table(), record.row_meta().primary_key(), &map)
    }

    pub fn hydrate<T>(record: &mut T, values: &RowMap) -> Result<(), DbError>
    where
        T: Record,
    {
        record.apply_row_map(values)
    }
}

pub mod data {
    pub use super::DbError;
    pub use super::Operator;
    pub use super::OrderClause;
    pub use super::OrderDirection;
    pub use super::QueryBuilder;
    pub use super::Record;
    pub use super::Row;
    pub use super::RowExecutor;
    pub use super::{set_default_store, DBStore, Filter, RowMap};
}

pub mod memory;
