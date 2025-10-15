<?php

namespace Com\Mh\Mimanjar\Domain\Data;

trait DiscountColumnsTraits
{
    public function code()
    {
        $this->addColumn( 'code' );
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

    public function description()
    {
        $this->addColumn( 'description' );
        return $this;
    }

    public function duration()
    {
        $this->addColumn( 'duration' );
        return $this;
    }

    public function endDate()
    {
        $this->addColumn( 'end_date' );
        return $this;
    }

    public function id()
    {
        $this->addColumn( 'id' );
        return $this;
    }

    public function locale()
    {
        $this->addColumn( 'locale' );
        return $this;
    }

    public function maxAmount()
    {
        $this->addColumn( 'max_amount' );
        return $this;
    }

    public function minAmount()
    {
        $this->addColumn( 'min_amount' );
        return $this;
    }

    public function name()
    {
        $this->addColumn( 'name' );
        return $this;
    }

    public function scope()
    {
        $this->addColumn( 'scope' );
        return $this;
    }

    public function startDate()
    {
        $this->addColumn( 'start_date' );
        return $this;
    }

    public function status()
    {
        $this->addColumn( 'status' );
        return $this;
    }

    public function target()
    {
        $this->addColumn( 'target' );
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

    public function value()
    {
        $this->addColumn( 'value' );
        return $this;
    }

}

