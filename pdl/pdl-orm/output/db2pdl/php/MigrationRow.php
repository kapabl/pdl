<?php
/**
 *  Minglehouse Generated File with Pdl
 *  MiManjar
 *  @version: 1.0.0
 *
 */

namespace Com\Mh\Mimanjar\Domain\Data;

use Com\Mh\Ds\Infrastructure\Data\Row;
/**
 * @class MigrationRow
 * @extends Row
 * @package Com\Mh\Mimanjar\Domain\Data
 *
 * @property int $batch
 * @property int $id
 * @property string $migration
 *
 */
class MigrationRow extends Row
{

    const TableName = 'migrations';
    const DbName = 'mimanjar_db';
    const FullTableName = 'mimanjar_db.migrations';

    /**
     * MigrationRow constructor.
     *
     * @param array $arguments
     */
    function __construct( ...$arguments )
    {
        $this->_propertyAttributes = [];
        $this->_propertyAttributes[ 'id' ] = [ 'IsDbId' => [] ];

        parent::__construct( ...$arguments );
    }

    /**
     * @return MigrationWhere
     */
    public static function where()
    {
        $result = new MigrationWhere( null );
        $result->rowClass = static::class;
        return $result;
    }

    /**
     * @return MigrationFieldList
     */
    public static function fieldList()
    {
        $result = new MigrationFieldList( null );
        $result->rowClass = static::class;
        return $result;
    }

}

