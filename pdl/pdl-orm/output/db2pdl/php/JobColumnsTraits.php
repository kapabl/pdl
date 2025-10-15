<?php

namespace Com\Mh\Mimanjar\Domain\Data;

trait JobColumnsTraits
{
    public function attempts()
    {
        $this->addColumn( 'attempts' );
        return $this;
    }

    public function availableAt()
    {
        $this->addColumn( 'available_at' );
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

    public function payload()
    {
        $this->addColumn( 'payload' );
        return $this;
    }

    public function queue()
    {
        $this->addColumn( 'queue' );
        return $this;
    }

    public function reservedAt()
    {
        $this->addColumn( 'reserved_at' );
        return $this;
    }

}

