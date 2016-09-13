package com.taijin.android.services;

import com.taijin.android.util.FoursquareResponse;

import retrofit2.Call;
import retrofit2.Retrofit;
import retrofit2.converter.gson.GsonConverterFactory;
import retrofit2.http.GET;
import retrofit2.http.Query;

/**
 * Created by taijin on 9/4/16.
 */
public interface FoursquareVenueInfoService {

    @GET("search")
    Call<FoursquareResponse> getRestaurantInfoFromFoursquare(@Query("client_id") String id,
                                                                   @Query("client_secret") String secret,
                                                                   @Query("ll") String lalong,
                                                                   @Query("v") String date);
    public static final Retrofit retrofit = new Retrofit.Builder()
            .baseUrl("https://api.foursquare.com/v2/venues/")
            .addConverterFactory(GsonConverterFactory.create())
            .build();
}
