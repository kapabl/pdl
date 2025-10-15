<?php

namespace Com\Mh\Mimanjar\Domain\Data;

trait StoreColumnsTraits
{
    public function businessHours()
    {
        $this->addColumn( 'business_hours' );
        return $this;
    }

    public function businessPhone()
    {
        $this->addColumn( 'business_phone' );
        return $this;
    }

    public function businessPhoto()
    {
        $this->addColumn( 'business_photo' );
        return $this;
    }

    public function categories()
    {
        $this->addColumn( 'categories' );
        return $this;
    }

    public function createdAt()
    {
        $this->addColumn( 'created_at' );
        return $this;
    }

    public function currency()
    {
        $this->addColumn( 'currency' );
        return $this;
    }

    public function datetimeZone()
    {
        $this->addColumn( 'datetime_zone' );
        return $this;
    }

    public function deliveryAddressId()
    {
        $this->addColumn( 'delivery_address_id' );
        return $this;
    }

    public function deliveryCost()
    {
        $this->addColumn( 'delivery_cost' );
        return $this;
    }

    public function deliveryOptions()
    {
        $this->addColumn( 'delivery_options' );
        return $this;
    }

    public function deliveryTimeframeOptions()
    {
        $this->addColumn( 'delivery_timeframe_options' );
        return $this;
    }

    public function id()
    {
        $this->addColumn( 'id' );
        return $this;
    }

    public function isAcceptingOrders()
    {
        $this->addColumn( 'is_accepting_orders' );
        return $this;
    }

    public function isOpen()
    {
        $this->addColumn( 'is_open' );
        return $this;
    }

    public function isPublished()
    {
        $this->addColumn( 'is_published' );
        return $this;
    }

    public function locale()
    {
        $this->addColumn( 'locale' );
        return $this;
    }

    public function name()
    {
        $this->addColumn( 'name' );
        return $this;
    }

    public function orderStateConfig()
    {
        $this->addColumn( 'order_state_config' );
        return $this;
    }

    public function paymentOptions()
    {
        $this->addColumn( 'payment_options' );
        return $this;
    }

    public function pickupAddressId()
    {
        $this->addColumn( 'pickup_address_id' );
        return $this;
    }

    public function serviceFeeOptions()
    {
        $this->addColumn( 'service_fee_options' );
        return $this;
    }

    public function slogan()
    {
        $this->addColumn( 'slogan' );
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

