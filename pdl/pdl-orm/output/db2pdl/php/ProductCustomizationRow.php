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
 * @class ProductCustomizationRow
 * @extends Row
 * @package Com\Mh\Mimanjar\Domain\Data
 *
 * @property string $createdAt
 * @property string $defaultValue
 * @property string $description
 * @property int $id
 * @property string $isOption
 * @property string $isSoldOut
 * @property int $maxQuantity
 * @property int $minQuantity
 * @property string $name
 * @property string $nutritionalInfo
 * @property int $pickupPrice
 * @property int $price
 * @property string $showInOrder
 * @property string $status
 * @property string $type
 * @property string $updatedAt
 * @property int $userId
 * @property string $uuid
 *
 */
class ProductCustomizationRow extends Row
{

    const TableName = 'product_customizations';
    const DbName = 'mimanjar_db';
    const FullTableName = 'mimanjar_db.product_customizations';

    /**
     * ProductCustomizationRow constructor.
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
     * @return ProductCustomizationWhere
     */
    public static function where()
    {
        $result = new ProductCustomizationWhere( null );
        $result->rowClass = static::class;
        return $result;
    }

    /**
     * @return ProductCustomizationFieldList
     */
    public static function fieldList()
    {
        $result = new ProductCustomizationFieldList( null );
        $result->rowClass = static::class;
        return $result;
    }

}

