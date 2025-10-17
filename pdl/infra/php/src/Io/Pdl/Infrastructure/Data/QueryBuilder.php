<?php

namespace Io\Pdl\Infrastructure\Data;

class QueryBuilder
{
    private array $filters = [];
    private array $projection = [];
    private array $orderings = [];
    private ?int $limit = null;
    private ?int $offset = null;

    public function __construct(
        private string $table,
        private ?DBStore $store
    ) {
    }

    public function withStore(?DBStore $store): self
    {
        $this->store = $store;
        return $this;
    }

    public function filter(string $column, string $operator, mixed $value): self
    {
        $this->filters[] = [$column, $operator, $value];
        return $this;
    }

    public function project(string ...$columns): self
    {
        $this->projection = array_merge($this->projection, $columns);
        return $this;
    }

    public function offset(int $value): self
    {
        $this->offset = $value;
        return $this;
    }

    public function limit(int $value): self
    {
        $this->limit = $value;
        return $this;
    }

    public function range(int $offset, int $limit): self
    {
        $this->offset = $offset;
        $this->limit = $limit;
        return $this;
    }

    public function orderBy(string $column, string $direction): self
    {
        $this->orderings[] = [$column, $direction];
        return $this;
    }

    public function load(): array
    {
        $store = StoreRegistry::resolve($this->store);
        return $store->select(
            $this->table,
            $this->filters,
            $this->projection,
            $this->orderings,
            $this->limit,
            $this->offset
        );
    }

    public function delete(string $primaryKey): void
    {
        $store = StoreRegistry::resolve($this->store);
        $rows = $store->select(
            $this->table,
            $this->filters,
            [$primaryKey],
            [],
            null,
            null
        );
        foreach ($rows as $row) {
            $store->delete($this->table, $primaryKey, $row);
        }
    }
}
