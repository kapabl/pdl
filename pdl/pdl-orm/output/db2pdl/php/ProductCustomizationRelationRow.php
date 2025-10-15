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
 * @class ProductCustomizationRelationRow
 * @extends Row
 * @package Com\Mh\Mimanjar\Domain\Data
 *
 * @property string $createdAt
 * @property int $id
 * @property int $parentId
 * @property int $position
 * @property int $productCustomizationId
 * @property int $productId
 * @property string $updatedAt
 * @property int $userId
 *
 */
class ProductCustomizationRelationRow extends Row
{

    const TableName = 'product_customization_relations';
    const DbName = 'mimanjar_db';
    const FullTableName = 'mimanjar_db.product_customization_relations';

    /**
     * ProductCustomizationRelationRow constructor.
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
     * @return ProductCustomizationRelationWhere
     */
    public static function where()
    {
        $result = new ProductCustomizationRelationWhere( null );
        $result->rowClass = static::class;
        return $result;
    }

    /**
     * @return ProductCustomizationRelationFieldList
     */
    public static function fieldList()
    {
        $result = new ProductCustomizationRelationFieldList( null );
        $result->rowClass = static::class;
        return $result;
    }

}

