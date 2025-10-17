<?php

namespace Io\Pdl\Infrastructure\Data;

interface DBStore
{
    public function insert(string $table, string $primaryKey, array $values): array;

    public function update(string $table, string $primaryKey, array $values): void;

    public function delete(string $table, string $primaryKey, array $values): void;

    public function select(
        string $table,
        array $filters,
        array $projection,
        array $orderings,
        ?int $limit,
        ?int $offset
    ): array;
}
