

```json
"db2pdl": {
  "enabled": true,
  "verbose": false,
  "connection": {
    "type": "mysql",
    "host": "${DB_HOST}",
    "port": "${DB_PORT}",
    "user": "${DB_USERNAME}",
    "password": "${DB_PASSWORD}",
    "database": "${DB_DATABASE}"
  },
  "outputDir": "output/pdl",
  "templatesDir": "default",
  "php": {
    "emitHelpers": true,
    "attributes": {
      "dbId": "[ 'IsDbId' => [] ]",
      "columnName": "[ 'ColumnName' => [ 'default1' => '{$column_name}' ] ]"
    }
  },
  "pdl": {
    "db2PdlSourceDest": "com/mh/mimanjar/domain/data",
    "entitiesNamespace": "com.mh.mimanjar.domain.data",
    "useNamespaces": [],
    "attributes": {
      "dbId": "[io.pdl.infrastructure.data.attributes.IsDbId]",
      "columnName": "[io.pdl.infrastructure.data.attributes.ColumnName(\"{$column_name}\")]"
    }
  },
  "excludedTables": [],
  "excludedColumns": []
}

```

## Runtime usage in `pdl-orm`

**Read by the Go db2pdl runner**

- `db2pdl.enabled`
- `db2pdl.verbose`
- `db2pdl.connection.type`
- `db2pdl.connection.host`
- `db2pdl.connection.port`
- `db2pdl.connection.user`
- `db2pdl.connection.password`
- `db2pdl.connection.database`
- `db2pdl.outputDir`
- `db2pdl.templatesDir`
- `db2pdl.ts.emit`
- `db2pdl.ts.outputFile`
- `db2pdl.cs.emit`
- `db2pdl.go.emit`
- `db2pdl.go.package`
- `db2pdl.php.emitHelpers`
- `db2pdl.php.attributes.dbId`
- `db2pdl.php.attributes.columnName`
- `db2pdl.pdl.entitiesNamespace`
  - Required; configuration load will panic if it is missing or empty
- `db2pdl.pdl.useNamespaces`
  - Alias: `db2pdl.pdl.use`
- `db2pdl.pdl.attributes.dbId`
- `db2pdl.pdl.attributes.columnName`
- `db2pdl.excludedTables`
- `db2pdl.excludedColumns`
