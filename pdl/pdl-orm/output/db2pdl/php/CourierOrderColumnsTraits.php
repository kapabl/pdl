<?php

namespace Com\Mh\Mimanjar\Domain\Data;

trait CourierOrderColumnsTraits
{
    public function courierUuid()
    {
        $this->addColumn( 'courier_uuid' );
        return $this;
    }

    public function createdAt()
    {
        $this->addColumn( 'created_at' );
        return $this;
    }

    public function id()
    {
        $this->addColumn( 'id' );
        return $this;
    }

    public function orderId()
    {
        $this->addColumn( 'order_id' );
        return $this;
    }

    public function status()
    {
        $this->addColumn( 'status' );
        return $this;
    }

    public function updatedAt()
    {
        $this->addColumn( 'updated_at' );
        return $this;
    }

}

