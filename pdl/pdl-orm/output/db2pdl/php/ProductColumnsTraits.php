<?php

namespace Com\Mh\Mimanjar\Domain\Data;

trait ProductColumnsTraits
{
    public function createdAt()
    {
        $this->addColumn( 'created_at' );
        return $this;
    }

    public function deliveryPrice()
    {
        $this->addColumn( 'delivery_price' );
        return $this;
    }

    public function depositOptions()
    {
        $this->addColumn( 'deposit_options' );
        return $this;
    }

    public function description()
    {
        $this->addColumn( 'description' );
        return $this;
    }

    public function details()
    {
        $this->addColumn( 'details' );
        return $this;
    }

    public function featured()
    {
        $this->addColumn( 'featured' );
        return $this;
    }

    public function id()
    {
        $this->addColumn( 'id' );
        return $this;
    }

    public function images()
    {
        $this->addColumn( 'images' );
        return $this;
    }

    public function keywords()
    {
        $this->addColumn( 'keywords' );
        return $this;
    }

    public function name()
    {
        $this->addColumn( 'name' );
        return $this;
    }

    public function price()
    {
        $this->addColumn( 'price' );
        return $this;
    }

    public function quantity()
    {
        $this->addColumn( 'quantity' );
        return $this;
    }

    public function slug()
    {
        $this->addColumn( 'slug' );
        return $this;
    }

    public function status()
    {
        $this->addColumn( 'status' );
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

