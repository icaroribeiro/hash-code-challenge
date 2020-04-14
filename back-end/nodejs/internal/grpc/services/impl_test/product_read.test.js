var setup = require('../../../tests/setup.js');
var grpc = require('grpc');

var ProductService = require('../../../grpc/services/product.js');

const mongoose = require('mongoose');

var Promotion = require('../../../models/promotion.js');
var DiscountedDate = require('../../../models/discounted-date.js');
var User = require('../../../models/user.js');
var Product = require('../../../models/product.js');

var Utils = require('../../../utils/index.js');

describe("TestGetAllProducts", () => {
    let grpcAddress;
    let client;
    let metadata;
    let currDate;

    let promotion;
    let product;
    let user = null
    let discountedDate1 = null
    let discountedDate2 = null

    beforeAll(async () => {
        client = new ProductService.stubConstructor(setup.GrpcAddress, grpc.credentials.createInsecure());
        
        metadata = new grpc.Metadata();

        currDate = Utils.GetCurrentDate();

        console.log("Configuring the data...");

        // Delete the promotion of discounted dates if it already exists.
        await Promotion.deleteOne({code: "DISCOUNTEDDATES"})
            .exec()
            .then(() => {
            })
            .catch(err => { 
                console.log(err);
                throw err;
            });

        promotion = new Promotion({
            _id: new mongoose.Types.ObjectId(),
            code: "DISCOUNTEDDATES",
            title: "Discounted Dates",
            description: "The promotion of discounted dates",
            max_discount_pct: 12.0
        });

        await promotion.save()
            .then((doc) => {
                console.log("Promotion:", JSON.stringify({
                    id: doc._id,
                    code: doc.code,
                    title: doc.title,
                    description: doc.description,
                    max_discount_pct: doc.max_discount_pct
                }, null, 0));
            })
            .catch(err => { 
                console.log(err);
                throw err;
            });

        product = new Product({
            _id: new mongoose.Types.ObjectId(),
            price_in_cents: 100,
            title: "Blue Pen",
            description: "A pen with blue ink"
        });

        setup.Datastore.CreateProduct(product, async function (err, data) {
            if (err) {
                console.log(err);
                throw err;
            }

            console.log("Product:", JSON.stringify({
                id: data.id,
                price_in_cents: data.price_in_cents,
                title: data.title,
                description: data.description
            }, null, 0));

            product._id = data.id;
        });
        
        /*await product.save()
            .then((doc) => {
                console.log("Product:", JSON.stringify({
                    id: doc._id,
                    price_in_cents: doc.price_in_cents,
                    title: doc.title,
                    description: doc.description
                }, null, 0));
            })
            .catch(err => { 
                console.log(err);
                throw err;
            });
        */

        promotion.products.push(product._id);

        await promotion.save()
            .then((doc) => {
                var products = [];
                doc.products.forEach((elem) => { 
                    products.push(elem);
                });

                console.log("New promotion data:", JSON.stringify({
                    id: doc._id,
                    code: doc.code,
                    title: doc.title,
                    description: doc.description,
                    max_discount_pct: doc.max_discount_pct,
                    products: products,
                }, null, 0));
            })
            .catch(err => { 
                console.log(err);
                throw err;
            });
    });

    afterAll(async () => {
    });

    afterEach(async () => {
    });

    test('WithoutAnyDiscountOfDates', async done => {
        await client.GetAllProducts({}, null, (err, response) => {
            if (!err) {  
                var mockData = {
                    id: `${product._id}`,
                    price_in_cents: product.price_in_cents,
                    title: product.title, 
                    description: product.description,
                    discount: null
                };

                response.products.forEach((elem) => {   
                    if (elem.id === mockData.id) {
                        console.log("Returned product:", JSON.stringify({
                            id: elem.id,
                            price_in_cents: elem.price_in_cents,
                            title: elem.title, 
                            description: elem.description,
                            discount: elem.discount
                        }, null, 0));

                        expect(elem).toEqual(mockData);
                    }
                });

                done();
            } else {
                console.log(err);
                done();
            }
        });
    });

    /*
    test('WithOnlyTheDiscountOfUser\'sBirthday', async done => {
    });

    test('WithOnlyTheDiscountOfOtherDiscountedDate', async done => {
    });

    test('WithTheMaximumDiscountOfDates', async done => {
    });
    */
});

/*
var grpc = require('grpc');

const { TEST_GRPC_SERVER_HOST, TEST_GRPC_SERVER_PORT } = require('../../../tests/env.js');

var ProductService = require('../../../grpc/services/product.js');

const mongoose = require('mongoose');

var Promotion = require('../../../models/promotion.js');
var DiscountedDate = require('../../../models/discounted-date.js');
var User = require('../../../models/user.js');
var Product = require('../../../models/product.js');

var Utils = require('../../../utils/index.js');

describe("TestGetAllProducts", () => {
    let grpcAddress;
    let client;
    let metadata;
    let currDate;

    let promotion;
    let product;
    let user = null
    let discountedDate1 = null
    let discountedDate2 = null

    beforeAll(async () => {
        try {
            var grpcHost = TEST_GRPC_SERVER_HOST;
        
            if (!grpcHost) {
                throw "Failed to read the TEST_GRPC_SERVER_HOST environment variable: it isn't set";
            }
        
            var grpcPort = TEST_GRPC_SERVER_PORT;
        
            if (!grpcPort) {
                throw "Failed to read the TEST_GRPC_SERVER_PORT environment variable: it isn't set";
            }
        }
        catch (err) {
            console.log(err);
            throw err;
        }
    
        grpcAddress = grpcHost + ":" + grpcPort;

        client = new ProductService.stubConstructor(grpcAddress, grpc.credentials.createInsecure());
        
        metadata = new grpc.Metadata();

        currDate = Utils.GetCurrentDate();

        console.log("Configuring the data...");

        // Delete the promotion of discounted dates if it already exists.
        await Promotion.deleteOne({code: "DISCOUNTEDDATES"})
            .exec()
            .then(() => {
            })
            .catch(err => { 
                console.log(err);
                throw err;
            });

        promotion = new Promotion({
            _id: new mongoose.Types.ObjectId(),
            code: "DISCOUNTEDDATES",
            title: "Discounted Dates",
            description: "The promotion of discounted dates",
            max_discount_pct: 12.0
        });

        await promotion.save()
            .then((doc) => {
                console.log("Promotion:", JSON.stringify({
                    id: doc._id,
                    code: doc.code,
                    title: doc.title,
                    description: doc.description,
                    max_discount_pct: doc.max_discount_pct
                }, null, 0));
            })
            .catch(err => { 
                console.log(err);
                throw err;
            });

        product = new Product({
            _id: new mongoose.Types.ObjectId(),
            price_in_cents: 100,
            title: "Blue Pen",
            description: "A pen with blue ink"
        });
        
        await product.save()
            .then((doc) => {
                console.log("Product:", JSON.stringify({
                    id: doc._id,
                    price_in_cents: doc.price_in_cents,
                    title: doc.title,
                    description: doc.description
                }, null, 0));
            })
            .catch(err => { 
                console.log(err);
                throw err;
            });

        promotion.products.push(product._id);

        await promotion.save()
            .then((doc) => {
                var products = [];
                doc.products.forEach((elem) => { 
                    products.push(elem);
                });

                console.log("Update promotion:", JSON.stringify({
                    id: doc._id,
                    code: doc.code,
                    title: doc.title,
                    description: doc.description,
                    max_discount_pct: doc.max_discount_pct,
                    products: products,
                }, null, 0));
            })
            .catch(err => { 
                console.log(err);
                throw err;
            });
    });

    afterAll(async () => {
        await Promotion.deleteOne({_id: promotion._id})
            .exec()
            .then(() => {
            })
            .catch(err => { 
                console.log(err);
                throw err;
            });

        await Product.deleteOne({_id: product._id})
            .exec()
            .then(() => {
            })
            .catch(err => { 
                console.log(err);
                throw err;
            });
    });

    afterEach(async () => {
        // Remove the X-USER-ID key and its associated value. 
        metadata.remove('X-USER-ID');

        if (user !== null) {
            await User.deleteOne({_id: user._id})
                .exec()
                .then(() => {
                })
                .catch(err => { 
                    console.log(err);
                    throw err;
                });
        }

        if (discountedDate1 !== null) {
            await DiscountedDate.deleteOne({_id: discountedDate1._id})
                .exec()
                .then(() => {
                })
                .catch(err => { 
                    console.log(err);
                    throw err;
                });
        }

        if (discountedDate2 !== null) {
            await DiscountedDate.deleteOne({_id: discountedDate2._id})
                .exec()
                .then(() => {
                })
                .catch(err => { 
                    console.log(err);
                    throw err;
                });
        }
    });

    test('WithoutAnyDiscountOfDates', async done => {        
        await client.GetAllProducts({}, null, (err, response) => {
            if (!err) {               
                var mockData = {
                    id: `${product._id}`,
                    price_in_cents: product.price_in_cents,
                    title: product.title, 
                    description: product.description,
                    discount: null
                };

                response.products.forEach((elem) => { 
                    if (elem.id === mockData.id) {
                        console.log("Returned product:", JSON.stringify({
                            id: elem.id,
                            price_in_cents: elem.price_in_cents,
                            title: elem.title, 
                            description: elem.description,
                            discount: elem.discount
                        }, null, 0));

                        expect(elem).toEqual(mockData);
                    }
                });

                done();
            } else {
                console.log(err);
                done();
            }
        });
    });

    test('WithOnlyTheDiscountOfUser\'sBirthday', async done => {
        var discountPct = 5.0
        var priceInCents = product.price_in_cents
        var valueInCents = (priceInCents * (discountPct / 100));
        priceInCents = priceInCents - Math.round(valueInCents);

        user = new User({
            _id: new mongoose.Types.ObjectId(),
            first_name: "User",
            last_name: "User",
            date_of_birth: {
                year: currDate.year,
                month: currDate.month,
                day: currDate.day,
            }
        });
        
        await user.save()
            .then((doc) => {
                console.log("User:", JSON.stringify({
                    id: doc._id,
                    first_name: doc.first_name,
                    last_name: doc.last_name,
                    date_of_birth: doc.date_of_birth
                }, null, 0));
            })
            .catch(err => { 
                console.log(err);
                throw err;
            });

        discountedDate1 = new DiscountedDate({
            _id: new mongoose.Types.ObjectId(),
            title: "User's Birthday",
            description: "The discount of user's birthday",
            discount_pct: discountPct,
            date: {
                year: 0,
                month: 0,
                day: 0,
            }
        });

        await discountedDate1.save()
            .then((doc) => {
                console.log("Discounted date:", JSON.stringify({
                    id: doc._id,
                    title: doc.title,
                    description: doc.description,
                    discount_pct: doc.discount_pct,
                    date: doc.date
                }, null, 0));
            })
            .catch(err => { 
                console.log(err);
                throw err;
            });

        metadata.add('X-USER-ID', `${user._id}`);
    
        await client.GetAllProducts({}, metadata, (err, response) => {
            if (!err) {            
                var discount = {
                    pct: discountPct, 
                    value_in_cents: {
                        value: valueInCents
                    }
                };

                var mockData = {
                    id: `${product._id}`,
                    price_in_cents: priceInCents,
                    title: product.title, 
                    description: product.description,
                    discount: discount
                };

                response.products.forEach((elem) => { 
                    if (elem.id === mockData.id) {
                        console.log("Returned product:", JSON.stringify({
                            id: elem.id,
                            price_in_cents: elem.price_in_cents,
                            title: elem.title, 
                            description: elem.description,
                            discount: elem.discount
                        }, null, 0));

                        expect(elem).toEqual(mockData);
                    }
                });

                done();
            } else {
                console.log(err);
                done();
            }
        });
    });

    test('WithOnlyTheDiscountOfOtherDiscountedDate', async done => {
        var discountPct = 10.0
        var priceInCents = product.price_in_cents
        var valueInCents = (priceInCents * (discountPct / 100));
        priceInCents = priceInCents - Math.round(valueInCents);
        
        user = new User({
            _id: new mongoose.Types.ObjectId(),
            first_name: "User",
            last_name: "User",
            date_of_birth: {
                year: currDate.year,
                month: currDate.month,
                day: currDate.day,
            }
        });
        
        await user.save()
            .then((doc) => {
                console.log("User:", JSON.stringify({
                    id: doc._id,
                    first_name: doc.first_name,
                    last_name: doc.last_name,
                    date_of_birth: doc.date_of_birth
                }, null, 0));
            })
            .catch(err => { 
                console.log(err);
                throw err;
            });

        discountedDate1 = new DiscountedDate({
            _id: new mongoose.Types.ObjectId(),
            title: "Other Discounted Date",
            description: "The discount of other discounted date",
            discount_pct: discountPct,
            date: {
                year: currDate.year,
                month: currDate.month,
                day: currDate.day,
            }
        });

        await discountedDate1.save()
            .then((doc) => {
                console.log("Other discounted date:", JSON.stringify({
                    id: doc._id,
                    title: doc.title,
                    discount_pct: doc.discount_pct,
                    date: doc.date
                }, null, 0));
            })
            .catch(err => { 
                console.log(err);
                throw err;
            });
            
        metadata.add('X-USER-ID', `${user._id}`);

        await client.GetAllProducts({}, metadata, (err, response) => {
            if (!err) {                               
                var discount = {
                    pct: discountPct, 
                    value_in_cents: {
                        value: valueInCents
                    }
                };

                var mockData = {
                    id: `${product._id}`,
                    price_in_cents: priceInCents,
                    title: product.title, 
                    description: product.description,
                    discount: discount
                };

                response.products.forEach((elem) => { 
                    if (elem.id === mockData.id) {
                        console.log("Returned product:", JSON.stringify({
                            id: elem.id,
                            price_in_cents: elem.price_in_cents,
                            title: elem.title, 
                            description: elem.description,
                            discount: elem.discount
                        }, null, 0));

                        expect(elem).toEqual(mockData);
                    }
                });

                done();
            } else {
                console.log(err);
                done();
            }
        });
    });

    test('WithTheMaximumDiscountOfDates', async done => {
        var discountPct1 = 5.0
        var discountPct2 = 10.0
        var maxDiscountPct = 12.0
        var priceInCents = product.price_in_cents
        var valueInCents = (priceInCents * (maxDiscountPct / 100));
        priceInCents = priceInCents - Math.round(valueInCents);

        user = new User({
            _id: new mongoose.Types.ObjectId(),
            first_name: "User",
            last_name: "User",
            date_of_birth: {
                year: currDate.year,
                month: currDate.month,
                day: currDate.day,
            }
        });
        
        await user.save()
            .then((doc) => {
                console.log("User:", JSON.stringify({
                    id: doc._id,
                    first_name: doc.first_name,
                    last_name: doc.last_name,
                    date_of_birth: doc.date_of_birth
                }, null, 0));
            })
            .catch(err => { 
                console.log(err);
                throw err;
            });

        discountedDate1 = new DiscountedDate({
            _id: new mongoose.Types.ObjectId(),
            title: "User's Birthday",
            description: "The discount of user's birthday",
            discount_pct: discountPct1,
            date: {
                year: 0,
                month: 0,
                day: 0,
            }
        });

        await discountedDate1.save()
            .then((doc) => {
                console.log("Discounted date:", JSON.stringify({
                    id: doc._id,
                    title: doc.title,
                    description: doc.description,
                    discount_pct: doc.discount_pct,
                    date: doc.date
                }, null, 0));
            })
            .catch(err => { 
                console.log(err);
                throw err;
            });

        discountedDate2 = new DiscountedDate({
            _id: new mongoose.Types.ObjectId(),
            title: "Other Discounted Date",
            description: "The discount of other discounted date",
            discount_pct: discountPct2,
            date: {
                year: currDate.year,
                month: currDate.month,
                day: currDate.day,
            }
        });

        await discountedDate2.save()
            .then((doc) => {
                console.log("Other discounted date:", JSON.stringify({
                    id: doc._id,
                    title: doc.title,
                    description: doc.description,
                    discount_pct: doc.discount_pct,
                    date: doc.date
                }, null, 0));
            })
            .catch(err => { 
                console.log(err);
                throw err;
            });
            
        metadata.add('X-USER-ID', `${user._id}`);

        await client.GetAllProducts({}, metadata, (err, response) => {
            if (!err) {                              
                var discount = {
                    pct: maxDiscountPct, 
                    value_in_cents: {
                        value: valueInCents
                    }
                };

                var mockData = {
                    id: `${product._id}`,
                    price_in_cents: priceInCents,
                    title: product.title, 
                    description: product.description,
                    discount: discount
                };

                response.products.forEach((elem) => {
                    if (elem.id === mockData.id) {
                        console.log("Returned product:", JSON.stringify({
                            id: elem.id,
                            price_in_cents: elem.price_in_cents,
                            title: elem.title, 
                            description: elem.description,
                            discount: elem.discount
                        }, null, 0));

                        expect(elem).toEqual(mockData);
                    }
                });

                done();
            } else {
                console.log(err);
                done();
            }
        });
    });
});

describe("TestGetProduct", () => {
    let grpcAddress;
    let client;
    let metadata;
    let currDate;

    let promotion;
    let product;
    let user = null
    let discountedDate1 = null
    let discountedDate2 = null

    beforeAll(async () => {
        try {
            var grpcHost = TEST_GRPC_SERVER_HOST;
        
            if (!grpcHost) {
                throw "Failed to read the TEST_GRPC_SERVER_HOST environment variable: it isn't set";
            }
        
            var grpcPort = TEST_GRPC_SERVER_PORT;
        
            if (!grpcPort) {
                throw "Failed to read the TEST_GRPC_SERVER_PORT environment variable: it isn't set";
            }
        }
        catch (err) {
            console.log(err);
            throw err;
        }

        grpcAddress = grpcHost + ":" + grpcPort;

        client = new ProductService.stubConstructor(grpcAddress, grpc.credentials.createInsecure());

        metadata = new grpc.Metadata();

        currDate = Utils.GetCurrentDate();

        console.log("Configuring the data...");

        // Delete the promotion of discounted dates if it already exists.
        await Promotion.deleteOne({code: "DISCOUNTEDDATES"})
            .exec()
            .then(() => {
            })
            .catch(err => { 
                console.log(err);
                throw err;
            });

        promotion = new Promotion({
            _id: new mongoose.Types.ObjectId(),
            code: "DISCOUNTEDDATES",
            title: "Discounted Dates",
            description: "The promotion of discounted dates",
            max_discount_pct: 12.0
        });

        await promotion.save()
            .then((doc) => {
                console.log("Promotion:", JSON.stringify({
                    id: doc._id,
                    code: doc.code,
                    title: doc.title,
                    description: doc.description,
                    max_discount_pct: doc.max_discount_pct
                }, null, 0));
            })
            .catch(err => { 
                console.log(err);
                throw err;
            });

        product = new Product({
            _id: new mongoose.Types.ObjectId(),
            price_in_cents: 100,
            title: "Blue Pen",
            description: "A pen with blue ink"
        });

        await product.save()
            .then((doc) => {
                console.log("Product:", JSON.stringify({
                    id: doc._id,
                    price_in_cents: doc.price_in_cents,
                    title: doc.title,
                    description: doc.description
                }, null, 0));
            })
            .catch(err => { 
                console.log(err);
                throw err;
            });

        promotion.products.push(product._id);

        await promotion.save()
            .then((doc) => {
                var products = [];
                doc.products.forEach((elem) => { 
                    products.push(elem);
                });

                console.log("Update promotion:", JSON.stringify({
                    id: doc._id,
                    code: doc.code,
                    title: doc.title,
                    description: doc.description,
                    max_discount_pct: doc.max_discount_pct,
                    products: products,
                }, null, 0));
            })
            .catch(err => { 
                console.log(err);
                throw err;
            });
    });

    afterAll(async () => {
        await Promotion.deleteOne({_id: promotion._id})
            .exec()
            .then(() => {
            })
            .catch(err => { 
                console.log(err);
                throw err;
            });

        await Product.deleteOne({_id: product._id})
            .exec()
            .then(() => {
            })
            .catch(err => { 
                console.log(err);
                throw err;
            });
    });

    afterEach(async () => {
        // Remove the X-USER-ID key and its associated value. 
        metadata.remove('X-USER-ID');

        if (user !== null) {
            await User.deleteOne({_id: user._id})
                .exec()
                .then(() => {
                })
                .catch(err => { 
                    console.log(err);
                    throw err;
                });
        }

        if (discountedDate1 !== null) {
            await DiscountedDate.deleteOne({_id: discountedDate1._id})
                .exec()
                .then(() => {
                })
                .catch(err => {
                    console.log(err);
                    throw err;
                });
        }

        if (discountedDate2 !== null) {
            await DiscountedDate.deleteOne({_id: discountedDate2._id})
                .exec()
                .then(() => {
                })
                .catch(err => {
                    console.log(err);
                    throw err;
                });
        }
    });

    test('WithoutAnyDiscountOfDates', async done => {
        await client.GetProduct({id: product._id}, null, (err, response) => {
            if (!err) {               
                var mockData = {
                    id: `${product._id}`,
                    price_in_cents: product.price_in_cents,
                    title: product.title, 
                    description: product.description,
                    discount: null
                };

                if (response.id === mockData.id) {
                    console.log("Returned product:", JSON.stringify({
                        id: response.id,
                        price_in_cents: response.price_in_cents,
                        title: response.title, 
                        description: response.description,
                        discount: response.discount
                    }, null, 0));

                    expect(response).toEqual(mockData);
                }

                done();
            } else {
                console.log(err);
                done();
            }
        });
    });

    test('WithOnlyTheDiscountOfUser\'sBirthday', async done => {
        var discountPct = 5.0
        var priceInCents = product.price_in_cents
        var valueInCents = (priceInCents * (discountPct / 100));
        priceInCents = priceInCents - Math.round(valueInCents);

        user = new User({
            _id: new mongoose.Types.ObjectId(),
            first_name: "User",
            last_name: "User",
            date_of_birth: {
                year: currDate.year,
                month: currDate.month,
                day: currDate.day,
            }
        });
        
        await user.save()
            .then((doc) => {
                console.log("User:", JSON.stringify({
                    id: doc._id,
                    first_name: doc.first_name,
                    last_name: doc.last_name,
                    date_of_birth: doc.date_of_birth
                }, null, 0));
            })
            .catch(err => { 
                console.log(err);
                throw err;
            });

        discountedDate1 = new DiscountedDate({
            _id: new mongoose.Types.ObjectId(),
            title: "User's Birthday",
            description: "The discount of user's birthday",
            discount_pct: discountPct,
            date: {
                year: 0,
                month: 0,
                day: 0,
            }
        });

        await discountedDate1.save()
            .then((doc) => {
                console.log("Discounted date:", JSON.stringify({
                    id: doc._id,
                    title: doc.title,
                    description: doc.description,
                    discount_pct: doc.discount_pct,
                    date: doc.date
                }, null, 0));
            })
            .catch(err => { 
                console.log(err);
                throw err;
            });

        metadata.add('X-USER-ID', `${user._id}`);
    
        await client.GetProduct({id: product._id}, metadata, (err, response) => {
            if (!err) {            
                var discount = {
                    pct: discountPct, 
                    value_in_cents: {
                        value: valueInCents
                    }
                };

                var mockData = {
                    id: `${product._id}`,
                    price_in_cents: priceInCents,
                    title: product.title, 
                    description: product.description,
                    discount: discount
                };

                if (response.id === mockData.id) {
                    console.log("Returned product:", JSON.stringify({
                        id: response.id,
                        price_in_cents: response.price_in_cents,
                        title: response.title, 
                        description: response.description,
                        discount: response.discount
                    }, null, 0));

                    expect(response).toEqual(mockData);
                }

                done();
            } else {
                console.log(err);
                done();
            }
        });
    });

    test('WithOnlyTheDiscountOfOtherDiscountedDate', async done => {
        var discountPct = 10.0
        var priceInCents = product.price_in_cents
        var valueInCents = (priceInCents * (discountPct / 100));
        priceInCents = priceInCents - Math.round(valueInCents);

        user = new User({
            _id: new mongoose.Types.ObjectId(),
            first_name: "User",
            last_name: "User",
            date_of_birth: {
                year: currDate.year,
                month: currDate.month,
                day: currDate.day,
            }
        });

        await user.save()
            .then((doc) => {
                console.log("User:", JSON.stringify({
                    id: doc._id,
                    first_name: doc.first_name,
                    last_name: doc.last_name,
                    date_of_birth: doc.date_of_birth
                }, null, 0));
            })
            .catch(err => { 
                console.log(err);
                throw err;
            });

        discountedDate1 = new DiscountedDate({
            _id: new mongoose.Types.ObjectId(),
            title: "Other Discounted Date",
            description: "The discount of other discounted date",
            discount_pct: discountPct,
            date: {
                year: currDate.year,
                month: currDate.month,
                day: currDate.day,
            }
        });

        await discountedDate1.save()
            .then((doc) => {
                console.log("Other discounted date:", JSON.stringify({
                    id: doc._id,
                    title: doc.title,
                    description: doc.description,
                    discount_pct: doc.discount_pct,
                    date: doc.date
                }, null, 0));
            })
            .catch(err => { 
                console.log(err);
                throw err;
            });

        metadata.add('X-USER-ID', `${user._id}`);

        await client.GetProduct({id: product._id}, metadata, (err, response) => {
            if (!err) {                               
                var discount = {
                    pct: discountPct, 
                    value_in_cents: {
                        value: valueInCents
                    }
                };

                var mockData = {
                    id: `${product._id}`,
                    price_in_cents: priceInCents,
                    title: product.title, 
                    description: product.description,
                    discount: discount
                };

                if (response.id === mockData.id) {
                    console.log("Returned product:", JSON.stringify({
                        id: response.id,
                        price_in_cents: response.price_in_cents,
                        title: response.title, 
                        description: response.description,
                        discount: response.discount
                    }, null, 0));

                    expect(response).toEqual(mockData);
                }

                done();
            } else {
                console.log(err);
                done();
            }
        });
    });

    test('WithTheMaximumDiscountOfDates', async done => {
        var discountPct1 = 5.0
        var discountPct2 = 10.0
        var maxDiscountPct = 12.0
        var priceInCents = product.price_in_cents
        var valueInCents = (priceInCents * (maxDiscountPct / 100));
        priceInCents = priceInCents - Math.round(valueInCents);

        user = new User({
            _id: new mongoose.Types.ObjectId(),
            first_name: "User",
            last_name: "User",
            date_of_birth: {
                year: currDate.year,
                month: currDate.month,
                day: currDate.day,
            }
        });

        await user.save()
            .then((doc) => {
                console.log("User:", JSON.stringify({
                    id: doc._id,
                    first_name: doc.first_name,
                    last_name: doc.last_name,
                    date_of_birth: doc.date_of_birth
                }, null, 0));
            })
            .catch(err => { 
                console.log(err);
                throw err;
            });

        discountedDate1 = new DiscountedDate({
            _id: new mongoose.Types.ObjectId(),
            title: "User's Birthday",
            description: "The discount of user's birthday",
            discount_pct: discountPct1,
            date: {
                year: 0,
                month: 0,
                day: 0,
            }
        });

        await discountedDate1.save()
            .then((doc) => {
                console.log("Discounted date:", JSON.stringify({
                    id: doc._id,
                    title: doc.title,
                    description: doc.description,
                    discount_pct: doc.discount_pct,
                    date: doc.date
                }, null, 0));
            })
            .catch(err => { 
                console.log(err);
                throw err;
            });

        discountedDate2 = new DiscountedDate({
            _id: new mongoose.Types.ObjectId(),
            title: "Other Discounted Date",
            description: "The discount of other discounted date",
            discount_pct: discountPct2,
            date: {
                year: currDate.year,
                month: currDate.month,
                day: currDate.day,
            }
        });

        await discountedDate2.save()
            .then((doc) => {
                console.log("Other discounted date:", JSON.stringify({
                    id: doc._id,
                    title: doc.title,
                    discount_pct: doc.discount_pct,
                    date: doc.date
                }, null, 0));
            })
            .catch(err => { 
                console.log(err);
                throw err;
            });

        metadata.add('X-USER-ID', `${user._id}`);

        await client.GetProduct({id: product._id}, metadata, (err, response) => {
            if (!err) {                              
                var discount = {
                    pct: maxDiscountPct, 
                    value_in_cents: {
                        value: valueInCents
                    }
                };

                var mockData = {
                    id: `${product._id}`,
                    price_in_cents: priceInCents,
                    title: product.title, 
                    description: product.description,
                    discount: discount
                };

                if (response.id === mockData.id) {
                    console.log("Returned product:", JSON.stringify({
                        id: response.id,
                        price_in_cents: response.price_in_cents,
                        title: response.title, 
                        description: response.description,
                        discount: response.discount
                    }, null, 0));

                    expect(response).toEqual(mockData);
                }

                done();
            } else {
                console.log(err);
                done();
            }
        });
    });
});
*/