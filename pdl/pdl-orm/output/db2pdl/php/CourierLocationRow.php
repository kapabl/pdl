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
 * @class CourierLocationRow
 * @extends Row
 * @package Com\Mh\Mimanjar\Domain\Data
 *
 * @property float $accuracy
 * @property float $altitude
 * @property float $altitudeAccuracy
 * @property string $courierUuid
 * @property float $heading
 * @property int $id
 * @property float $lat
 * @property string $locationTime
 * @property float $lon
 * @property float $speed
 *
 */
class CourierLocationRow extends Row
{

    const TableName = 'courier_locations';
    const DbName = 'mimanjar_db';
    const FullTableName = 'mimanjar_db.courier_locations';

    /**
     * CourierLocationRow constructor.
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
     * @return CourierLocationWhere
     */
    public static function where()
    {
        $result = new CourierLocationWhere( null );
        $result->rowClass = static::class;
        return $result;
    }

    /**
     * @return CourierLocationFieldList
     */
    public static function fieldList()
    {
        $result = new CourierLocationFieldList( null );
        $result->rowClass = static::class;
        return $result;
    }

}

