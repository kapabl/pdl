<?php

namespace Com\Mh\Mimanjar\Domain\Data;

trait MenuColumnsTraits
{
    public function categoryIds()
    {
        $this->addColumn( 'category_ids' );
        return $this;
    }

    public function createdAt()
    {
        $this->addColumn( 'created_at' );
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

    public function name()
    {
        $this->addColumn( 'name' );
        return $this;
    }

    public function schedule()
    {
        $this->addColumn( 'schedule' );
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

}

