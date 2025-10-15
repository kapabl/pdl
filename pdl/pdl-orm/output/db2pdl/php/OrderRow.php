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
 * @class OrderRow
 * @extends Row
 * @package Com\Mh\Mimanjar\Domain\Data
 *
 * @property string $courierUuid
 * @property string $createdAt
 * @property int $deliveryTaxAmount
 * @property int $discountAmount
 * @property int $id
 * @property string $model
 * @property string $notes
 * @property string $orderId
 * @property int $sellerTaxAmount
 * @property int $sellerTotal
 * @property string $status
 * @property int $storeId
 * @property int $subTotal
 * @property int $taxAmount
 * @property int $total
 * @property string $updatedAt
 * @property int $userId
 *
 */
class OrderRow extends Row
{

    const TableName = 'orders';
    const DbName = 'mimanjar_db';
    const FullTableName = 'mimanjar_db.orders';

    /**
     * OrderRow constructor.
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
     * @return OrderWhere
     */
    public static function where()
    {
        $result = new OrderWhere( null );
        $result->rowClass = static::class;
        return $result;
    }

    /**
     * @return OrderFieldList
     */
    public static function fieldList()
    {
        $result = new OrderFieldList( null );
        $result->rowClass = static::class;
        return $result;
    }

}

