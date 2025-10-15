<?php

namespace Com\Mh\Mimanjar\Domain\Data;

trait DiscountInstanceColumnsTraits
{
    public function createdAt()
    {
        $this->addColumn( 'created_at' );
        return $this;
    }

    public function discountId()
    {
        $this->addColumn( 'discount_id' );
        return $this;
    }

    public function id()
    {
        $this->addColumn( 'id' );
        return $this;
    }

    public function instanceInfo()
    {
        $this->addColumn( 'instance_info' );
        return $this;
    }

    public function productId()
    {
        $this->addColumn( 'product_id' );
        return $this;
    }

    public function storeId()
    {
        $this->addColumn( 'store_id' );
        return $this;
    }

    public function updatedAt()
    {
        $this->addColumn( 'updated_at' );
        return $this;
    }

}

