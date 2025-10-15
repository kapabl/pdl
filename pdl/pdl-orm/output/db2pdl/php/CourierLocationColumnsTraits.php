<?php

namespace Com\Mh\Mimanjar\Domain\Data;

trait CourierLocationColumnsTraits
{
    public function accuracy()
    {
        $this->addColumn( 'accuracy' );
        return $this;
    }

    public function altitude()
    {
        $this->addColumn( 'altitude' );
        return $this;
    }

    public function altitudeAccuracy()
    {
        $this->addColumn( 'altitude_accuracy' );
        return $this;
    }

    public function courierUuid()
    {
        $this->addColumn( 'courier_uuid' );
        return $this;
    }

    public function heading()
    {
        $this->addColumn( 'heading' );
        return $this;
    }

    public function id()
    {
        $this->addColumn( 'id' );
        return $this;
    }

    public function lat()
    {
        $this->addColumn( 'lat' );
        return $this;
    }

    public function locationTime()
    {
        $this->addColumn( 'location_time' );
        return $this;
    }

    public function lon()
    {
        $this->addColumn( 'lon' );
        return $this;
    }

    public function speed()
    {
        $this->addColumn( 'speed' );
        return $this;
    }

}

