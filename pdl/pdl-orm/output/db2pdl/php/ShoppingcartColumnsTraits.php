<?php

namespace Com\Mh\Mimanjar\Domain\Data;

trait ShoppingcartColumnsTraits
{
    public function content()
    {
        $this->addColumn( 'content' );
        return $this;
    }

    public function createdAt()
    {
        $this->addColumn( 'created_at' );
        return $this;
    }

    public function identifier()
    {
        $this->addColumn( 'identifier' );
        return $this;
    }

    public function instance()
    {
        $this->addColumn( 'instance' );
        return $this;
    }

    public function updatedAt()
    {
        $this->addColumn( 'updated_at' );
        return $this;
    }

}

