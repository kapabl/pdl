<?php

namespace Com\Mh\Mimanjar\Domain\Data;

trait PersonalAccessTokenColumnsTraits
{
    public function abilities()
    {
        $this->addColumn( 'abilities' );
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

    public function lastUsedAt()
    {
        $this->addColumn( 'last_used_at' );
        return $this;
    }

    public function name()
    {
        $this->addColumn( 'name' );
        return $this;
    }

    public function token()
    {
        $this->addColumn( 'token' );
        return $this;
    }

    public function tokenableId()
    {
        $this->addColumn( 'tokenable_id' );
        return $this;
    }

    public function tokenableType()
    {
        $this->addColumn( 'tokenable_type' );
        return $this;
    }

    public function updatedAt()
    {
        $this->addColumn( 'updated_at' );
        return $this;
    }

}

