<?php

namespace Com\Mh\Mimanjar\Domain\Data;

trait TeamUserColumnsTraits
{
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

    public function role()
    {
        $this->addColumn( 'role' );
        return $this;
    }

    public function teamId()
    {
        $this->addColumn( 'team_id' );
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

}

