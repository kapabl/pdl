<?php

namespace Com\Mh\Mimanjar\Domain\Data;

use Com\Mh\Ds\Infrastructure\Data\Db\ColumnsDefinition;

class NotificationColumns extends ColumnsDefinition
{
    protected function setup()
    {
        $this->addColumn( 'body', 'string' );
        $this->addColumn( 'created_at', 'string' );
        $this->addColumn( 'data', 'string' );
        $this->addColumn( 'from_store_Id', 'int' );
        $this->addColumn( 'from_user_id', 'int' );
        $this->addColumn( 'id', 'int' );
        $this->addColumn( 'name', 'string' );
        $this->addColumn( 'notification_time', 'string' );
        $this->addColumn( 'order_id', 'string' );
        $this->addColumn( 'read_time', 'string' );
        $this->addColumn( 'retrieved_first_time', 'string' );
        $this->addColumn( 'source', 'string' );
        $this->addColumn( 'title', 'string' );
        $this->addColumn( 'to_store_id', 'int' );
        $this->addColumn( 'to_user_id', 'int' );
        $this->addColumn( 'type', 'string' );
        $this->addColumn( 'unread', 'string' );
        $this->addColumn( 'updated_at', 'string' );
        $this->addColumn( 'url', 'string' );
    }
}

