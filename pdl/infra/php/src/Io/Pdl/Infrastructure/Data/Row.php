<?php

namespace Io\Pdl\Infrastructure\Data;

use RuntimeException;

class Row
{
    private array $propertyToColumn;
    private array $columnToProperty;
    private ?DBStore $store;

    public function __construct(
        private string $table,
        private string $primaryKey,
        array $columnMap,
        ?DBStore $store = null
    ) {
        $this->propertyToColumn = $columnMap;
        $this->columnToProperty = array_change_key_case(array_flip($columnMap));
        $this->store = $store;
    }

    public function table(): string
    {
        return $this->table;
    }

    public function primaryKey(): string
    {
        return $this->primaryKey;
    }

    public function store(): ?DBStore
    {
        return $this->store;
    }

    public function setStore(?DBStore $store): void
    {
        $this->store = $store;
    }

    public function collectValues(): array
    {
        $values = [];
        foreach ($this->propertyToColumn as $property => $column) {
            if (!property_exists($this, $property)) {
                continue;
            }
            $value = $this->{$property};
            $values[$column] = $value;
        }
        return $values;
    }

    public function applyValues(array $values): void
    {
        foreach ($values as $column => $value) {
            $normalized = strtolower($column);
            $property = $this->columnToProperty[$normalized] ?? $this->derivePropertyName($column);
            if ($property === null || !property_exists($this, $property)) {
                continue;
            }
            $this->{$property} = $value;
        }
    }

    private function derivePropertyName(string $column): ?string
    {
        $normalized = strtolower($column);
        if (isset($this->columnToProperty[$normalized])) {
            return $this->columnToProperty[$normalized];
        }
        $segments = explode('_', $normalized);
        $segments = array_map(static fn(string $part): string => ucfirst($part), $segments);
        $property = lcfirst(implode('', $segments));
        return property_exists($this, $property) ? $property : null;
    }

    public function resolveStore(): DBStore
    {
        return StoreRegistry::resolve($this->store);
    }
}
