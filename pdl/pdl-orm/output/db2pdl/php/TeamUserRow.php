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
 * @class TeamUserRow
 * @extends Row
 * @package Com\Mh\Mimanjar\Domain\Data
 *
 * @property string $createdAt
 * @property int $id
 * @property string $role
 * @property int $teamId
 * @property string $updatedAt
 * @property int $userId
 *
 */
class TeamUserRow extends Row
{

    const TableName = 'team_user';
    const DbName = 'mimanjar_db';
    const FullTableName = 'mimanjar_db.team_user';

    /**
     * TeamUserRow constructor.
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
     * @return TeamUserWhere
     */
    public static function where()
    {
        $result = new TeamUserWhere( null );
        $result->rowClass = static::class;
        return $result;
    }

    /**
     * @return TeamUserFieldList
     */
    public static function fieldList()
    {
        $result = new TeamUserFieldList( null );
        $result->rowClass = static::class;
        return $result;
    }

}

