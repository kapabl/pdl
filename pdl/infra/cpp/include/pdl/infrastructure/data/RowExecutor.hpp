#pragma once

#include <memory>

#include "pdl/infrastructure/data/DBStore.hpp"

namespace pdl::infrastructure::data {

void SetDefaultStore(DBStorePtr store);

DBStorePtr DefaultStore();

namespace detail {

DBStorePtr ResolveStore(const DBStorePtr& overrideStore);

} // namespace detail

class RowExecutor {
public:
    template <typename Record>
    static void Create(Record& record) {
        auto values = record.CollectValues();
        auto store = detail::ResolveStore(record.RowMetadata().Store());
        auto inserted = store->Insert(record.RowMetadata().Table(), record.RowMetadata().PrimaryKey(), std::move(values));
        record.RowMetadata().SetStore(store);
        if (!inserted.empty()) {
            record.ApplyValues(inserted);
        }
    }

    template <typename Record>
    static void Update(Record& record) {
        auto values = record.CollectValues();
        auto store = detail::ResolveStore(record.RowMetadata().Store());
        record.RowMetadata().SetStore(store);
        store->Update(record.RowMetadata().Table(), record.RowMetadata().PrimaryKey(), values);
    }

    template <typename Record>
    static void Delete(Record& record) {
        auto values = record.CollectValues();
        auto store = detail::ResolveStore(record.RowMetadata().Store());
        record.RowMetadata().SetStore(store);
        store->Delete(record.RowMetadata().Table(), record.RowMetadata().PrimaryKey(), values);
    }

    template <typename Record>
    static void Hydrate(Record& record, const RowMap& values) {
        record.ApplyValues(values);
    }
};

} // namespace pdl::infrastructure::data
