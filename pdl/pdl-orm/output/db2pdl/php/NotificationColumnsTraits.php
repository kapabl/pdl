<?php

namespace Com\Mh\Mimanjar\Domain\Data;

trait NotificationColumnsTraits
{
    public function body()
    {
        $this->addColumn( 'body' );
        return $this;
    }

    public function createdAt()
    {
        $this->addColumn( 'created_at' );
        return $this;
    }

    public function data()
    {
        $this->addColumn( 'data' );
        return $this;
    }

    public function fromStoreId()
    {
        $this->addColumn( 'from_store_Id' );
        return $this;
    }

    public function fromUserId()
    {
        $this->addColumn( 'from_user_id' );
        return $this;
    }

    public function id()
    {
        $this->addColumn( 'id' );
        return $this;
    }

    public function name()
    {
        $this->addColumn( 'name' );
        return $this;
    }

    public function notificationTime()
    {
        $this->addColumn( 'notification_time' );
        return $this;
    }

    public function orderId()
    {
        $this->addColumn( 'order_id' );
        return $this;
    }

    public function readTime()
    {
        $this->addColumn( 'read_time' );
        return $this;
    }

    public function retrievedFirstTime()
    {
        $this->addColumn( 'retrieved_first_time' );
        return $this;
    }

    public function source()
    {
        $this->addColumn( 'source' );
        return $this;
    }

    public function title()
    {
        $this->addColumn( 'title' );
        return $this;
    }

    public function toStoreId()
    {
        $this->addColumn( 'to_store_id' );
        return $this;
    }

    public function toUserId()
    {
        $this->addColumn( 'to_user_id' );
        return $this;
    }

    public function type()
    {
        $this->addColumn( 'type' );
        return $this;
    }

    public function unread()
    {
        $this->addColumn( 'unread' );
        return $this;
    }

    public function updatedAt()
    {
        $this->addColumn( 'updated_at' );
        return $this;
    }

    public function url()
    {
        $this->addColumn( 'url' );
        return $this;
    }

}

