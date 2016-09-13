package com.taijin.android.util;

import com.google.gson.annotations.SerializedName;

/**
 * Created by taijin on 9/4/16.
 */
public class FoursquareResponse {
    @SerializedName("response")
    public FoursquareVenueList list;

    public FoursquareResponse(FoursquareVenueList list) {
        this.list = list;
    }

    public FoursquareVenueList getList() {
        return list;
    }

    public void setList(FoursquareVenueList list) {
        this.list = list;
    }
}
