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
 * @class NotificationRow
 * @extends Row
 * @package Com\Mh\Mimanjar\Domain\Data
 *
 * @property string $body
 * @property string $createdAt
 * @property string $data
 * @property int $fromStoreId
 * @property int $fromUserId
 * @property int $id
 * @property string $name
 * @property string $notificationTime
 * @property string $orderId
 * @property string $readTime
 * @property string $retrievedFirstTime
 * @property string $source
 * @property string $title
 * @property int $toStoreId
 * @property int $toUserId
 * @property string $type
 * @property string $unread
 * @property string $updatedAt
 * @property string $url
 *
 */
class NotificationRow extends Row
{

    const TableName = 'notifications';
    const DbName = 'mimanjar_db';
    const FullTableName = 'mimanjar_db.notifications';

    /**
     * NotificationRow constructor.
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
     * @return NotificationWhere
     */
    public static function where()
    {
        $result = new NotificationWhere( null );
        $result->rowClass = static::class;
        return $result;
    }

    /**
     * @return NotificationFieldList
     */
    public static function fieldList()
    {
        $result = new NotificationFieldList( null );
        $result->rowClass = static::class;
        return $result;
    }

}

