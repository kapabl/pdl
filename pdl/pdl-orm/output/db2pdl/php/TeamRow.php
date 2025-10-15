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
 * @class TeamRow
 * @extends Row
 * @package Com\Mh\Mimanjar\Domain\Data
 *
 * @property string $createdAt
 * @property int $id
 * @property string $name
 * @property int $personalTeam
 * @property string $updatedAt
 * @property int $userId
 *
 */
class TeamRow extends Row
{

    const TableName = 'teams';
    const DbName = 'mimanjar_db';
    const FullTableName = 'mimanjar_db.teams';

    /**
     * TeamRow constructor.
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
     * @return TeamWhere
     */
    public static function where()
    {
        $result = new TeamWhere( null );
        $result->rowClass = static::class;
        return $result;
    }

    /**
     * @return TeamFieldList
     */
    public static function fieldList()
    {
        $result = new TeamFieldList( null );
        $result->rowClass = static::class;
        return $result;
    }

}

