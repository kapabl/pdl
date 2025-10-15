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
 * @class OrderProductRow
 * @extends Row
 * @package Com\Mh\Mimanjar\Domain\Data
 *
 * @property string $createdAt
 * @property int $id
 * @property int $orderId
 * @property int $productId
 * @property int $quantity
 * @property string $updatedAt
 *
 */
class OrderProductRow extends Row
{

    const TableName = 'order_product';
    const DbName = 'mimanjar_db';
    const FullTableName = 'mimanjar_db.order_product';

    /**
     * OrderProductRow constructor.
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
     * @return OrderProductWhere
     */
    public static function where()
    {
        $result = new OrderProductWhere( null );
        $result->rowClass = static::class;
        return $result;
    }

    /**
     * @return OrderProductFieldList
     */
    public static function fieldList()
    {
        $result = new OrderProductFieldList( null );
        $result->rowClass = static::class;
        return $result;
    }

}

