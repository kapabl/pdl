<?php

namespace Io\Pdl\Infrastructure\Data;

final class RowExecutor
{
    public static function create(Row $record): void
    {
        $values = $record->collectValues();
        $store = $record->resolveStore();
        $record->setStore($store);
        $inserted = $store->insert($record->table(), $record->primaryKey(), $values);
        if (!empty($inserted)) {
            $record->applyValues($inserted);
        }
    }

    public static function update(Row $record): void
    {
        $values = $record->collectValues();
        $store = $record->resolveStore();
        $record->setStore($store);
        $store->update($record->table(), $record->primaryKey(), $values);
    }

    public static function delete(Row $record): void
    {
        $values = $record->collectValues();
        $store = $record->resolveStore();
        $record->setStore($store);
        $store->delete($record->table(), $record->primaryKey(), $values);
    }

    public static function hydrate(Row $record, array $values): void
    {
        $record->applyValues($values);
    }
}
