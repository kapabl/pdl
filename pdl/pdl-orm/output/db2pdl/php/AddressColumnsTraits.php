<?php

namespace Com\Mh\Mimanjar\Domain\Data;

trait AddressColumnsTraits
{
    public function address1()
    {
        $this->addColumn( 'address1' );
        return $this;
    }

    public function address2()
    {
        $this->addColumn( 'address2' );
        return $this;
    }

    public function city()
    {
        $this->addColumn( 'city' );
        return $this;
    }

    public function country()
    {
        $this->addColumn( 'country' );
        return $this;
    }

    public function createdAt()
    {
        $this->addColumn( 'created_at' );
        return $this;
    }

    public function defaultDelivery()
    {
        $this->addColumn( 'default_delivery' );
        return $this;
    }

    public function defaultPickup()
    {
        $this->addColumn( 'default_pickup' );
        return $this;
    }

    public function id()
    {
        $this->addColumn( 'id' );
        return $this;
    }

    public function isTest()
    {
        $this->addColumn( 'is_test' );
        return $this;
    }

    public function lat()
    {
        $this->addColumn( 'lat' );
        return $this;
    }

    public function lon()
    {
        $this->addColumn( 'lon' );
        return $this;
    }

    public function name()
    {
        $this->addColumn( 'name' );
        return $this;
    }

    public function phone()
    {
        $this->addColumn( 'phone' );
        return $this;
    }

    public function state()
    {
        $this->addColumn( 'state' );
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

    public function userId()
    {
        $this->addColumn( 'user_id' );
        return $this;
    }

    public function zipcode()
    {
        $this->addColumn( 'zipcode' );
        return $this;
    }

}

