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
 * @class SessionRow
 * @extends Row
 * @package Com\Mh\Mimanjar\Domain\Data
 *
 * @property string $id
 * @property string $ipAddress
 * @property int $lastActivity
 * @property string $payload
 * @property string $userAgent
 * @property int $userId
 *
 */
class SessionRow extends Row
{

    const TableName = 'sessions';
    const DbName = 'mimanjar_db';
    const FullTableName = 'mimanjar_db.sessions';

    /**
     * SessionRow constructor.
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
     * @return SessionWhere
     */
    public static function where()
    {
        $result = new SessionWhere( null );
        $result->rowClass = static::class;
        return $result;
    }

    /**
     * @return SessionFieldList
     */
    public static function fieldList()
    {
        $result = new SessionFieldList( null );
        $result->rowClass = static::class;
        return $result;
    }

}

