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
 * @class ProductRow
 * @extends Row
 * @package Com\Mh\Mimanjar\Domain\Data
 *
 * @property string $createdAt
 * @property int $deliveryPrice
 * @property string $depositOptions
 * @property string $description
 * @property string $details
 * @property int $featured
 * @property int $id
 * @property string $images
 * @property string $keywords
 * @property string $name
 * @property int $price
 * @property int $quantity
 * @property string $slug
 * @property string $status
 * @property int $storeId
 * @property string $updatedAt
 * @property int $userId
 * @property string $uuid
 *
 */
class ProductRow extends Row
{

    const TableName = 'products';
    const DbName = 'mimanjar_db';
    const FullTableName = 'mimanjar_db.products';

    /**
     * ProductRow constructor.
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
     * @return ProductWhere
     */
    public static function where()
    {
        $result = new ProductWhere( null );
        $result->rowClass = static::class;
        return $result;
    }

    /**
     * @return ProductFieldList
     */
    public static function fieldList()
    {
        $result = new ProductFieldList( null );
        $result->rowClass = static::class;
        return $result;
    }

}

