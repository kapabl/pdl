#include "pdl/infrastructure/data/RowExecutor.hpp"

#include <mutex>
#include <stdexcept>

namespace pdl::infrastructure::data {

namespace {
std::mutex g_storeMutex;
DBStorePtr g_defaultStore;
}

void SetDefaultStore(DBStorePtr store) {
    std::lock_guard<std::mutex> lock(g_storeMutex);
    g_defaultStore = std::move(store);
}

DBStorePtr DefaultStore() {
    std::lock_guard<std::mutex> lock(g_storeMutex);
    return g_defaultStore;
}

namespace detail {

DBStorePtr ResolveStore(const DBStorePtr& overrideStore) {
    if (overrideStore) {
        return overrideStore;
    }
    auto store = DefaultStore();
    if (!store) {
        throw std::runtime_error("pdl::infrastructure::data: no default DBStore configured");
    }
    return store;
}

} // namespace detail

} // namespace pdl::infrastructure::data
