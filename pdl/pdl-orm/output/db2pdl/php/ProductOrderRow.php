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
 * @class ProductOrderRow
 * @extends Row
 * @package Com\Mh\Mimanjar\Domain\Data
 *
 * @property string $createdAt
 * @property int $id
 * @property int $position
 * @property int $productId
 * @property string $updatedAt
 * @property int $userId
 *
 */
class ProductOrderRow extends Row
{

    const TableName = 'product_order';
    const DbName = 'mimanjar_db';
    const FullTableName = 'mimanjar_db.product_order';

    /**
     * ProductOrderRow constructor.
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
     * @return ProductOrderWhere
     */
    public static function where()
    {
        $result = new ProductOrderWhere( null );
        $result->rowClass = static::class;
        return $result;
    }

    /**
     * @return ProductOrderFieldList
     */
    public static function fieldList()
    {
        $result = new ProductOrderFieldList( null );
        $result->rowClass = static::class;
        return $result;
    }

}

