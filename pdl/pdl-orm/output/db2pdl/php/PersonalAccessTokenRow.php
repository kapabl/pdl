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
 * @class PersonalAccessTokenRow
 * @extends Row
 * @package Com\Mh\Mimanjar\Domain\Data
 *
 * @property string $abilities
 * @property string $createdAt
 * @property int $id
 * @property string $lastUsedAt
 * @property string $name
 * @property string $token
 * @property int $tokenableId
 * @property string $tokenableType
 * @property string $updatedAt
 *
 */
class PersonalAccessTokenRow extends Row
{

    const TableName = 'personal_access_tokens';
    const DbName = 'mimanjar_db';
    const FullTableName = 'mimanjar_db.personal_access_tokens';

    /**
     * PersonalAccessTokenRow constructor.
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
     * @return PersonalAccessTokenWhere
     */
    public static function where()
    {
        $result = new PersonalAccessTokenWhere( null );
        $result->rowClass = static::class;
        return $result;
    }

    /**
     * @return PersonalAccessTokenFieldList
     */
    public static function fieldList()
    {
        $result = new PersonalAccessTokenFieldList( null );
        $result->rowClass = static::class;
        return $result;
    }

}

