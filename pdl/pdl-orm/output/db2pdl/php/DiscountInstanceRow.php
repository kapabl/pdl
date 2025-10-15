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
 * @class DiscountInstanceRow
 * @extends Row
 * @package Com\Mh\Mimanjar\Domain\Data
 *
 * @property string $createdAt
 * @property int $discountId
 * @property int $id
 * @property string $instanceInfo
 * @property int $productId
 * @property int $storeId
 * @property string $updatedAt
 *
 */
class DiscountInstanceRow extends Row
{

    const TableName = 'discount_instances';
    const DbName = 'mimanjar_db';
    const FullTableName = 'mimanjar_db.discount_instances';

    /**
     * DiscountInstanceRow constructor.
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
     * @return DiscountInstanceWhere
     */
    public static function where()
    {
        $result = new DiscountInstanceWhere( null );
        $result->rowClass = static::class;
        return $result;
    }

    /**
     * @return DiscountInstanceFieldList
     */
    public static function fieldList()
    {
        $result = new DiscountInstanceFieldList( null );
        $result->rowClass = static::class;
        return $result;
    }

}

