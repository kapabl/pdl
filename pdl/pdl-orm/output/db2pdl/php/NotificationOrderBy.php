<?php

namespace Com\Mh\Mimanjar\Domain\Data;

use Com\Mh\Ds\Infrastructure\Data\Db\OrderBy;

class NotificationOrderBy extends OrderBy
{
    protected function setup()
    {
        $this->addField( 'body' );
        $this->addField( 'created_at' );
        $this->addField( 'data' );
        $this->addField( 'from_store_Id' );
        $this->addField( 'from_user_id' );
        $this->addField( 'id' );
        $this->addField( 'name' );
        $this->addField( 'notification_time' );
        $this->addField( 'order_id' );
        $this->addField( 'read_time' );
        $this->addField( 'retrieved_first_time' );
        $this->addField( 'source' );
        $this->addField( 'title' );
        $this->addField( 'to_store_id' );
        $this->addField( 'to_user_id' );
        $this->addField( 'type' );
        $this->addField( 'unread' );
        $this->addField( 'updated_at' );
        $this->addField( 'url' );
    }
}

