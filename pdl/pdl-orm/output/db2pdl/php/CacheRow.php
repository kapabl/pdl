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
 * @class CacheRow
 * @extends Row
 * @package Com\Mh\Mimanjar\Domain\Data
 *
 * @property int $expiration
 * @property string $key
 * @property string $value
 *
 */
class CacheRow extends Row
{

    const TableName = 'cache';
    const DbName = 'mimanjar_db';
    const FullTableName = 'mimanjar_db.cache';

    /**
     * CacheRow constructor.
     *
     * @param array $arguments
     */
    function __construct( ...$arguments )
    {
        $this->_propertyAttributes = [];
        $this->_propertyAttributes[ 'key' ] = [ 'IsDbId' => [] ];

        parent::__construct( ...$arguments );
    }

    /**
     * @return CacheWhere
     */
    public static function where()
    {
        $result = new CacheWhere( null );
        $result->rowClass = static::class;
        return $result;
    }

    /**
     * @return CacheFieldList
     */
    public static function fieldList()
    {
        $result = new CacheFieldList( null );
        $result->rowClass = static::class;
        return $result;
    }

}

