const Promotion = require('../../models/promotion.js');

var DiscountedDates = require('./discounted-dates/index.js');

exports.EvaluatePromotions = async function(data, xUserId = "") {
    // A map which holds key-value pairs composed by the product's id and its related discount percentage
    // according to the settings of the promotions in which each product is included.
    var productsMap = new Map();

    await ExamineDiscountedDates(productsMap, xUserId);

    // In case it is analyzed a single product.
    if (data.products === undefined) {
        if (productsMap.has(`${data.id}`)) {
            AddDiscountFieldToProduct(data, productsMap.get(`${data.id}`));
        }
    } else {
        // In case it is analyzed a list of one or more products.
        data.products.forEach((product) => {
            if (productsMap.has(`${product.id}`)) {
                AddDiscountFieldToProduct(product, productsMap.get(`${product.id}`));
            }
        });
    }
}

async function ExamineDiscountedDates(productsMap, xUserId) {
    var code = "DISCOUNTEDDATES"
    var data = {maxDiscountPct: 0, products: []};
    var discountPct = 0;

    await GetPromotionData(code, data);

    discountPct = await DiscountedDates.GetDiscountPct(xUserId);

    if (discountPct == 0) {
        return;
    }

    OrganizeProductDiscounts(productsMap, data, discountPct);
}

async function GetPromotionData(code, data) {
    await Promotion.findOne({code: code})
        .select('max_discount_pct products')
        .exec()
        .then((doc) => {
            if (doc) {
                data.maxDiscountPct = doc.max_discount_pct;
                
                if (doc.products !== undefined) {
                    doc.products.forEach((id) => {
                        data.products.push(id);
                    });
                }
            }
        })
        .catch(err => { 
            console.log(err);
            throw err;
        });
}

function OrganizeProductDiscounts(productsMap, data, discountPct) {
    // Limit the discount percentage of the products.
    if (data.maxDiscountPct == 0 || data.products.length == 0) {
        return;
    } else {
        if (discountPct > data.maxDiscountPct) {
            discountPct = data.maxDiscountPct;
        }
    }

    if (data.products !== undefined) {
        data.products.forEach((id) => {    
            if (productsMap.has(`${id}`)) {
                // In case a product is included in more than one promotion
                // it's necessary to aggregate the discount percentages.
                productsMap.set(id, productsMap.get(`${id}`) + discountPct);
            } else {
                productsMap.set(id, discountPct);
            }
        });
    }
}

function AddDiscountFieldToProduct(product, discountPct) {
    var priceInCents = 0
    var valueInCents = 0;

    priceInCents = product.price_in_cents;

    valueInCents = (priceInCents * (discountPct / 100));

    product.price_in_cents = priceInCents - Math.round(valueInCents);

    product.discount = {
        pct: discountPct,
        value_in_cents: {
            value: Math.round(valueInCents)
        }
    };
}