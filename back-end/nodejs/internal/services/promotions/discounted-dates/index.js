const User = require('../../../models/user.js');
const DiscountedDate = require('../../../models/discounted-date.js');

var Utils = require('../../../utils/index.js');

exports.GetDiscountPct = async function(xUserId) {
    var currDate = Utils.GetCurrentDate();    
    var discountPct = 0;

    // First, check if exists other special date that is different from user's birthday.
    await DiscountedDate.findOne({
            "date.year": currDate.year,
            "date.month": currDate.month,
            "date.day": currDate.day
        })
        .select('discount_pct')
        .exec()
        .then((doc) => {
            if (doc) {
                discountPct = discountPct + doc.discount_pct
            }
        })
        .catch(err => { 
            console.log(err);
            throw err;
        });

    // Second, analyze if there is any discount for user's birthday.
    if (xUserId !== "") {
        var dateOfBirth = null;

        await User.findOne({_id: xUserId})
            .select('date_of_birth')
            .exec()
            .then((doc) => {
                if (doc) {
                    dateOfBirth = doc.date_of_birth
                }
            })
            .catch(err => { 
                console.log(err);
                throw err;
            });

        if (dateOfBirth !== null) {            
            // In order to analyze the date of birth, consider only the day and month.
            delete currDate.year;
            delete dateOfBirth.year;

            // If the dates are equivalent, it means that today is the user's birth.
            // Therefore, seach for the discount date configured for that.
            if (Utils.IsEquivalent(currDate, dateOfBirth)) {
                await DiscountedDate.findOne({
                        "date.year": 0,
                        "date.month": 0,
                        "date.day": 0
                    })
                    .select('discount_pct')
                    .exec()
                    .then((doc) => {
                        if (doc) {
                            discountPct = discountPct + doc.discount_pct
                        }
                    })
                    .catch(err => { 
                        console.log(err);
                        throw err;
                    });
            }
        }
    }

    return discountPct;
}