package com.taijin.android;

import android.os.Bundle;
import android.support.v7.app.AppCompatActivity;

import com.taijin.android.dataRetrievingObjects.DataGettingResInfo;
import com.taijin.android.serverResponseObjects.AllTime;
import com.taijin.android.serverResponseObjects.Dinner;
import com.taijin.android.serverResponseObjects.Lunch;
import com.taijin.android.serverResponseObjects.RestaurantInfo;
import com.taijin.android.serverResponseObjects.Special;
import com.taijin.android.services.RestaurantInfoService;

import retrofit2.Call;

/**
 * Created by taijin on 8/29/16.
 */
public class CheckinActivity extends AppCompatActivity {

    private RestaurantInfo restaurantInfo;
    private Lunch lunch;
    private Dinner dinner;
    private AllTime alltime;
    private Special special;
    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        //setContentView(R.layout.activity_checkin);
        getMenu();

    }

    public void getMenu() {
        RestaurantInfoService restaurantInfoService = RestaurantInfoService.retrofit.create(RestaurantInfoService.class);
        final Call<RestaurantInfo> call =
                restaurantInfoService.getRestaurantInfo(new DataGettingResInfo("12345678", "taijin"), "getRestaurantInfo");
/*
        call.enqueue(new Callback<RestaurantInfo>() {
            @Override
            public void onResponse(Call<RestaurantInfo> call, Response<RestaurantInfo> response) {
                System.out.println("Response status code: " + response.code());

                final TextView textView = (TextView) findViewById(R.id.textView);
                response.errorBody();
                System.out.println(response.body().toString());
                textView.setText(response.body().toString());
            }
            @Override
            public void onFailure(Call<RestaurantInfo> call, Throwable t) {
                final TextView textView = (TextView) findViewById(R.id.textView);
                textView.setText("Something went wrong: " + t.getMessage());
            }
        });
        */
    }


}
