use std::collections::HashMap;
use std::sync::{Arc, Mutex};

use crate::{DbError, DBStore, Filter, Operator, OrderClause, OrderDirection, RowMap};
use serde_json::json;

#[derive(Default)]
struct Table {
    rows: Vec<RowMap>,
}

#[derive(Default)]
pub struct MemoryStore {
    tables: Mutex<HashMap<String, Table>>,
}

impl MemoryStore {
    pub fn new() -> Self {
        Self::default()
    }
}

impl DBStore for MemoryStore {
    fn insert(&self, table: &str, primary_key: &str, values: &mut RowMap) -> Result<RowMap, DbError> {
        let mut tables = self.tables.lock().expect("MemoryStore lock poisoned");
        let table_entry = tables.entry(table.to_string()).or_default();
        if !values.contains_key(primary_key) {
            let id = table_entry.rows.len() as i64 + 1;
            values.insert(primary_key.to_string(), serde_json::json!(id));
        }
        table_entry.rows.push(values.clone());
        Ok(values.clone())
    }

    fn update(&self, table: &str, primary_key: &str, values: &RowMap) -> Result<(), DbError> {
        let mut tables = self.tables.lock().expect("MemoryStore lock poisoned");
        let Some(table_entry) = tables.get_mut(table) else {
            return Err(DbError::new("MemoryStore: table not found"));
        };
        let key = values.get(primary_key).cloned();
        for row in &mut table_entry.rows {
            if row.get(primary_key) == key.as_ref() {
                for (column, value) in values {
                    row.insert(column.clone(), value.clone());
                }
                return Ok(());
            }
        }
        Err(DbError::new("MemoryStore: row not found"))
    }

    fn delete(&self, table: &str, primary_key: &str, values: &RowMap) -> Result<(), DbError> {
        let mut tables = self.tables.lock().expect("MemoryStore lock poisoned");
        let Some(table_entry) = tables.get_mut(table) else {
            return Ok(());
        };
        let key = values.get(primary_key).cloned();
        table_entry
            .rows
            .retain(|row| row.get(primary_key) != key.as_ref());
        Ok(())
    }

    fn select(
        &self,
        table: &str,
        filters: &[Filter],
        projection: &[String],
        orderings: &[OrderClause],
        limit: Option<i32>,
        offset: Option<i32>,
    ) -> Result<Vec<RowMap>, DbError> {
        let tables = self.tables.lock().expect("MemoryStore lock poisoned");
        let mut rows = tables
            .get(table)
            .map(|entry| entry.rows.clone())
            .unwrap_or_default();

        for filter in filters {
            rows.retain(|row| match filter.op {
                Operator::Eq => row.get(&filter.field) == Some(&filter.value),
                Operator::Neq => row.get(&filter.field) != Some(&filter.value),
                Operator::In => false,
            });
        }

        if !orderings.is_empty() {
            rows.sort_by(|left, right| {
                for order in orderings {
                    let comparison = left.get(&order.column).cmp(&right.get(&order.column));
                    if comparison != std::cmp::Ordering::Equal {
                        return match order.direction {
                            OrderDirection::Asc => comparison,
                            OrderDirection::Desc => comparison.reverse(),
                        };
                    }
                }
                std::cmp::Ordering::Equal
            });
        }

        let start = offset.unwrap_or(0).max(0) as usize;
        let mut end = limit.map(|value| start + value.max(0) as usize).unwrap_or(rows.len());
        end = end.min(rows.len());
        let mut sliced = rows.into_iter().skip(start).take(end.saturating_sub(start)).collect::<Vec<_>>();

        if !projection.is_empty() {
            sliced = sliced
                .into_iter()
                .map(|mut row| {
                    row.retain(|key, _| projection.contains(key));
                    row
                })
                .collect();
        }

        Ok(sliced)
    }
}

pub fn memory_store() -> Arc<dyn DBStore> {
    Arc::new(MemoryStore::new())
}
