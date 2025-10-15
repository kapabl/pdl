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
 * @class CategoryProductRow
 * @extends Row
 * @package Com\Mh\Mimanjar\Domain\Data
 *
 * @property int $categoryId
 * @property string $createdAt
 * @property int $id
 * @property int $position
 * @property int $productId
 * @property int $storeId
 * @property string $updatedAt
 *
 */
class CategoryProductRow extends Row
{

    const TableName = 'category_products';
    const DbName = 'mimanjar_db';
    const FullTableName = 'mimanjar_db.category_products';

    /**
     * CategoryProductRow constructor.
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
     * @return CategoryProductWhere
     */
    public static function where()
    {
        $result = new CategoryProductWhere( null );
        $result->rowClass = static::class;
        return $result;
    }

    /**
     * @return CategoryProductFieldList
     */
    public static function fieldList()
    {
        $result = new CategoryProductFieldList( null );
        $result->rowClass = static::class;
        return $result;
    }

}

