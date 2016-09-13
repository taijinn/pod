package com.taijin.android.util;

import com.google.gson.annotations.SerializedName;

import java.util.List;

/**
 * Created by taijin on 9/4/16.
 */
public class FoursquareVenueList {
    @SerializedName("venues")
    public List<FoursquareVenue> venues;

    public FoursquareVenueList(List<FoursquareVenue> venues) {
        this.venues = venues;
    }

    public List<FoursquareVenue> getVenues() {
        return venues;
    }

    public void setVenues(List<FoursquareVenue> venues) {
        this.venues = venues;
    }
}
