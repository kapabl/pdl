<?php

namespace Com\Mh\Mimanjar\Domain\Data;

trait OrderProductColumnsTraits
{
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

    public function productId()
    {
        $this->addColumn( 'product_id' );
        return $this;
    }

    public function quantity()
    {
        $this->addColumn( 'quantity' );
        return $this;
    }

    public function updatedAt()
    {
        $this->addColumn( 'updated_at' );
        return $this;
    }

}

