<?php

namespace Io\Pdl\Infrastructure\Data\Store;

use Io\Pdl\Infrastructure\Data\DBStore;
use Io\Pdl\Infrastructure\Data\Operator;
use RuntimeException;

final class MemoryStore implements DBStore
{
    /** @var array<string, array<int, array<string, mixed>>> */
    private array $tables = [];

    public function insert(string $table, string $primaryKey, array $values): array
    {
        $tableData = &$this->tables[$table];
        if ($tableData === null) {
            $tableData = [];
        }
        if (!array_key_exists($primaryKey, $values) || $values[$primaryKey] === null) {
            $values[$primaryKey] = count($tableData) + 1;
        }
        $tableData[] = $values;
        return [$primaryKey => $values[$primaryKey]];
    }

    public function update(string $table, string $primaryKey, array $values): void
    {
        foreach ($this->tables[$table] ?? [] as &$row) {
            if (($row[$primaryKey] ?? null) === ($values[$primaryKey] ?? null)) {
                $row = array_replace($row, $values);
                return;
            }
        }
        throw new RuntimeException('MemoryStore: row not found for update');
    }

    public function delete(string $table, string $primaryKey, array $values): void
    {
        foreach ($this->tables[$table] ?? [] as $index => $row) {
            if (($row[$primaryKey] ?? null) === ($values[$primaryKey] ?? null)) {
                unset($this->tables[$table][$index]);
                return;
            }
        }
    }

    public function select(
        string $table,
        array $filters,
        array $projection,
        array $orderings,
        ?int $limit,
        ?int $offset
    ): array {
        $rows = array_values($this->tables[$table] ?? []);
        foreach ($filters as [$column, $operator, $value]) {
            $rows = array_values(array_filter($rows, static function (array $row) use ($column, $operator, $value): bool {
                $current = $row[$column] ?? null;
                return match ($operator) {
                    Operator::EQ => $current === $value,
                    Operator::NEQ => $current !== $value,
                    Operator::IN => is_array($value) && in_array($current, $value, true),
                    default => false,
                };
            }));
        }
        if (!empty($orderings)) {
            usort($rows, static function (array $left, array $right) use ($orderings): int {
                foreach ($orderings as [$column, $direction]) {
                    $comparison = ($left[$column] ?? null) <=> ($right[$column] ?? null);
                    if ($comparison !== 0) {
                        return $direction === 'DESC' ? -$comparison : $comparison;
                    }
                }
                return 0;
            });
        }
        if ($offset !== null) {
            $rows = array_slice($rows, $offset);
        }
        if ($limit !== null) {
            $rows = array_slice($rows, 0, $limit);
        }
        if (!empty($projection)) {
            $rows = array_map(static function (array $row) use ($projection): array {
                $projected = [];
                foreach ($projection as $column) {
                    $projected[$column] = $row[$column] ?? null;
                }
                return $projected;
            }, $rows);
        }
        return $rows;
    }
}
