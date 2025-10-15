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
 * @class DiscountRow
 * @extends Row
 * @package Com\Mh\Mimanjar\Domain\Data
 *
 * @property string $code
 * @property string $createdAt
 * @property string $currency
 * @property string $datetimeZone
 * @property string $description
 * @property int $duration
 * @property string $endDate
 * @property int $id
 * @property string $locale
 * @property float $maxAmount
 * @property float $minAmount
 * @property string $name
 * @property string $scope
 * @property string $startDate
 * @property string $status
 * @property string $target
 * @property string $type
 * @property string $updatedAt
 * @property float $value
 *
 */
class DiscountRow extends Row
{

    const TableName = 'discounts';
    const DbName = 'mimanjar_db';
    const FullTableName = 'mimanjar_db.discounts';

    /**
     * DiscountRow constructor.
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
     * @return DiscountWhere
     */
    public static function where()
    {
        $result = new DiscountWhere( null );
        $result->rowClass = static::class;
        return $result;
    }

    /**
     * @return DiscountFieldList
     */
    public static function fieldList()
    {
        $result = new DiscountFieldList( null );
        $result->rowClass = static::class;
        return $result;
    }

}

