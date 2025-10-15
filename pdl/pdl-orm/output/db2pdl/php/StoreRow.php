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
 * @class StoreRow
 * @extends Row
 * @package Com\Mh\Mimanjar\Domain\Data
 *
 * @property string $businessHours
 * @property string $businessPhone
 * @property string $businessPhoto
 * @property string $categories
 * @property string $createdAt
 * @property string $currency
 * @property string $datetimeZone
 * @property int $deliveryAddressId
 * @property int $deliveryCost
 * @property string $deliveryOptions
 * @property string $deliveryTimeframeOptions
 * @property int $id
 * @property string $isAcceptingOrders
 * @property string $isOpen
 * @property string $isPublished
 * @property string $locale
 * @property string $name
 * @property string $orderStateConfig
 * @property string $paymentOptions
 * @property int $pickupAddressId
 * @property string $serviceFeeOptions
 * @property string $slogan
 * @property string $slug
 * @property string $status
 * @property string $updatedAt
 * @property int $userId
 * @property string $uuid
 *
 */
class StoreRow extends Row
{

    const TableName = 'stores';
    const DbName = 'mimanjar_db';
    const FullTableName = 'mimanjar_db.stores';

    /**
     * StoreRow constructor.
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
     * @return StoreWhere
     */
    public static function where()
    {
        $result = new StoreWhere( null );
        $result->rowClass = static::class;
        return $result;
    }

    /**
     * @return StoreFieldList
     */
    public static function fieldList()
    {
        $result = new StoreFieldList( null );
        $result->rowClass = static::class;
        return $result;
    }

}

