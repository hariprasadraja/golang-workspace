db.getCollection('system.js').save(
   {
     _id: "makeSubcategoryName",
     value : function() {
       var totalAccounts = 0
       var totalStore = 0
       var totalSaleUpdate = 0
       var categoryMap = [];
       var totalCategory = 0

       db.account.find({}).forEach(function(account) {
           totalAccounts = totalAccounts + 1;
        db.store.find({"ownedBy":account._id}).forEach(function(store){
                 totalStore = totalStore + 1;
          db.category.find({
                     "ownedBy": account._id,
                     "createdFor": store._id,
                     "level": 1
                     }).forEach(function(category){
                        categoryMap[category._id.toString()] = category.name
                         totalCategory = totalCategory + 1
                     })

          db.sale.find({
                    "account": account._id,
                    "store": store._id,
                    "orders": {
                            $elemMatch: {
                                "subCategoryName": {
                                 $exists: false
                                 },
                                 "categoryLevel": 2,
                              }
                        }
                  }, {
                    "_id": 1,
                    "orders": 1
              }).forEach(function(sale){
                  for (var i in sale.orders) {
              sale.orders[i].subCategoryName = sale.orders[i].categoryName;
              catID = sale.orders[i].category;
              sale.orders[i].categoryName = categoryMap[catID];
              db.sale.update(sale);
              totalSaleUpdate = totalSaleUpdate + 1
                  }
              })
           }) // store
      }) // account
              return {
                  "totalAccounts":totalAccounts,
                  "totalStore":totalStore,
                  "totalCategory":totalCategory,
                  "categoryMap":categoryMap,
                  "totalSaleUpdate":totalSaleUpdate,
                  }
     }
 })