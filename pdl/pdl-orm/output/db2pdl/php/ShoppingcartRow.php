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
 * @class ShoppingcartRow
 * @extends Row
 * @package Com\Mh\Mimanjar\Domain\Data
 *
 * @property string $content
 * @property string $createdAt
 * @property string $identifier
 * @property string $instance
 * @property string $updatedAt
 *
 */
class ShoppingcartRow extends Row
{

    const TableName = 'shoppingcart';
    const DbName = 'mimanjar_db';
    const FullTableName = 'mimanjar_db.shoppingcart';

    /**
     * ShoppingcartRow constructor.
     *
     * @param array $arguments
     */
    function __construct( ...$arguments )
    {
        $this->_propertyAttributes = [];
        $this->_propertyAttributes[ 'identifier' ] = [ 'IsDbId' => [] ];
        $this->_propertyAttributes[ 'instance' ] = [ 'IsDbId' => [] ];

        parent::__construct( ...$arguments );
    }

    /**
     * @return ShoppingcartWhere
     */
    public static function where()
    {
        $result = new ShoppingcartWhere( null );
        $result->rowClass = static::class;
        return $result;
    }

    /**
     * @return ShoppingcartFieldList
     */
    public static function fieldList()
    {
        $result = new ShoppingcartFieldList( null );
        $result->rowClass = static::class;
        return $result;
    }

}

