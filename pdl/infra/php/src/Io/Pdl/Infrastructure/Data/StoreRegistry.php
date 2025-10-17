<?php

namespace Io\Pdl\Infrastructure\Data;

use RuntimeException;

final class StoreRegistry
{
    private static ?DBStore $defaultStore = null;

    public static function setDefaultStore(DBStore $store): void
    {
        self::$defaultStore = $store;
    }

    public static function defaultStore(): DBStore
    {
        if (self::$defaultStore === null) {
            throw new RuntimeException('Io\Pdl\Infrastructure\Data: default store is not configured');
        }

        return self::$defaultStore;
    }

    public static function resolve(?DBStore $store): DBStore
    {
        return $store ?? self::defaultStore();
    }
}
