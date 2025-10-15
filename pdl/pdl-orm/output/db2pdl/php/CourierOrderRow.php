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
 * @class CourierOrderRow
 * @extends Row
 * @package Com\Mh\Mimanjar\Domain\Data
 *
 * @property string $courierUuid
 * @property string $createdAt
 * @property int $id
 * @property int $orderId
 * @property string $status
 * @property string $updatedAt
 *
 */
class CourierOrderRow extends Row
{

    const TableName = 'courier_orders';
    const DbName = 'mimanjar_db';
    const FullTableName = 'mimanjar_db.courier_orders';

    /**
     * CourierOrderRow constructor.
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
     * @return CourierOrderWhere
     */
    public static function where()
    {
        $result = new CourierOrderWhere( null );
        $result->rowClass = static::class;
        return $result;
    }

    /**
     * @return CourierOrderFieldList
     */
    public static function fieldList()
    {
        $result = new CourierOrderFieldList( null );
        $result->rowClass = static::class;
        return $result;
    }

}

