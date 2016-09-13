package com.taijin.android.services;

import com.taijin.android.util.GoogleVenueResult;

import retrofit2.Call;
import retrofit2.Retrofit;
import retrofit2.converter.gson.GsonConverterFactory;
import retrofit2.http.GET;
import retrofit2.http.Query;

/**
 * Created by taijin on 9/6/16.
 */
public interface GoogleVenueInfoService {
    @GET("json")
     Call<GoogleVenueResult> getRestaurantInfoFromGoogle(@Query("key") String key,
                                                         @Query("radius") int radius,
                                                         @Query("location") String lalong,
                                                         @Query("type") String date);
    public static final Retrofit retrofit = new Retrofit.Builder()
            .baseUrl("https://maps.googleapis.com/maps/api/place/nearbysearch/")
            .addConverterFactory(GsonConverterFactory.create())
            .build();
}
