<?php

namespace Com\Mh\Mimanjar\Domain\Data;

trait ProductCustomizationColumnsTraits
{
    public function createdAt()
    {
        $this->addColumn( 'created_at' );
        return $this;
    }

    public function defaultValue()
    {
        $this->addColumn( 'default_value' );
        return $this;
    }

    public function description()
    {
        $this->addColumn( 'description' );
        return $this;
    }

    public function id()
    {
        $this->addColumn( 'id' );
        return $this;
    }

    public function isOption()
    {
        $this->addColumn( 'is_option' );
        return $this;
    }

    public function isSoldOut()
    {
        $this->addColumn( 'is_sold_out' );
        return $this;
    }

    public function maxQuantity()
    {
        $this->addColumn( 'max_quantity' );
        return $this;
    }

    public function minQuantity()
    {
        $this->addColumn( 'min_quantity' );
        return $this;
    }

    public function name()
    {
        $this->addColumn( 'name' );
        return $this;
    }

    public function nutritionalInfo()
    {
        $this->addColumn( 'nutritional_info' );
        return $this;
    }

    public function pickupPrice()
    {
        $this->addColumn( 'pickup_price' );
        return $this;
    }

    public function price()
    {
        $this->addColumn( 'price' );
        return $this;
    }

    public function showInOrder()
    {
        $this->addColumn( 'show_in_order' );
        return $this;
    }

    public function status()
    {
        $this->addColumn( 'status' );
        return $this;
    }

    public function type()
    {
        $this->addColumn( 'type' );
        return $this;
    }

    public function updatedAt()
    {
        $this->addColumn( 'updated_at' );
        return $this;
    }

    public function userId()
    {
        $this->addColumn( 'user_id' );
        return $this;
    }

    public function uuid()
    {
        $this->addColumn( 'uuid' );
        return $this;
    }

}

