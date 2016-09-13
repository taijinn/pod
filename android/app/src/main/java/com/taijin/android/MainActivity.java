package com.taijin.android;

import android.content.Intent;
import android.os.Bundle;
import android.support.design.widget.FloatingActionButton;
import android.support.design.widget.Snackbar;
import android.support.v4.view.MenuItemCompat;
import android.support.v7.app.AppCompatActivity;
import android.support.v7.widget.SearchView;
import android.support.v7.widget.Toolbar;
import android.view.Menu;
import android.view.MenuItem;
import android.view.View;
import android.widget.AdapterView;
import android.widget.Button;
import android.widget.ListView;
import android.widget.TextView;
import android.widget.Toast;

import com.taijin.android.services.GoogleVenueInfoService;
import com.taijin.android.util.GoogleVenue;
import com.taijin.android.util.GoogleVenueResult;

import java.util.List;

import butterknife.BindView;
import butterknife.ButterKnife;
import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Response;

public class MainActivity extends AppCompatActivity implements SearchView.OnQueryTextListener {

    @BindView(R.id.listView)
    ListView lv;

    List<GoogleVenue> gv;

    @BindView(R.id.button_checkin)
    Button button_checkin;

    final String KEY = "AIzaSyBnfBG4lIb-PBNjnlMRFP3lpxsR6Wd-Jy4";
    final int radius = 500;
    //final String CLIENT_SECRET = "44SB1VXAWWNWOTNYBRQO30F5QSN4RABWP2AGOASQH3JFYO2Y";

    final String lalong = "37.377688,-122.055386";


    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
        ButterKnife.bind(this);

        Toolbar toolbar = (Toolbar) findViewById(R.id.toolbar);
        setSupportActionBar(toolbar);

        FloatingActionButton fab = (FloatingActionButton) findViewById(R.id.fab);
        fab.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View view) {
                Snackbar.make(view, "Super fast hello world from JRebel for Android", Snackbar.LENGTH_LONG)
                        .setAction("Action", null).show();
            }
        });


        GoogleVenueInfoService googleVenueInfoService = GoogleVenueInfoService.retrofit.create(GoogleVenueInfoService.class);

        Call<GoogleVenueResult> call =
                googleVenueInfoService.getRestaurantInfoFromGoogle(KEY, radius, lalong, "restaurant");

        call.enqueue(new Callback<GoogleVenueResult>() {
            @Override
            public void onResponse(Call<GoogleVenueResult> call, Response<GoogleVenueResult> response) {
                System.out.println("Response status code: " + response.code());
                if (response.isSuccessful()){
                    System.out.println(response.body().toString());
                    CustomAdapter adapter = new CustomAdapter(getBaseContext(), response.body().getVenues());
                    lv.setAdapter(adapter);
                }else{
                    System.out.println("not successfull.");
                }
            }
            @Override
            public void onFailure(Call<GoogleVenueResult> call, Throwable throwable) {
                System.out.println(throwable.getMessage());
            }

        });

        lv.setOnItemClickListener(listHandler);
        button_checkin.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View view) {

                Intent intent = new Intent(String.valueOf(CheckinActivity.class));

            }
        });


    }
    private AdapterView.OnItemClickListener listHandler = new AdapterView.OnItemClickListener() {
        @Override
        public void onItemClick(AdapterView<?> adapterView, View view, int i, long l) {
            TextView textView = (TextView) view.findViewById(R.id.listText);
            String res = textView.getText().toString();
            Toast.makeText(MainActivity.this, res, Toast.LENGTH_SHORT).show();

            Intent intent = new Intent(getApplicationContext(), VisitRestaurantActivity.class);
            intent.putExtra("restaurantName", res);
            startActivity(intent);

        }
    };
    @Override
    public boolean onCreateOptionsMenu(Menu menu) {
        // Inflate the menu; this adds items to the action bar if it is present.
        getMenuInflater().inflate(R.menu.menu_main, menu);
        MenuItem searchItem = menu.findItem(R.id.search);
        SearchView searchView = (SearchView) MenuItemCompat.getActionView(searchItem);
        searchView.setOnQueryTextListener(this);

        return true;
    }

    @Override
    public boolean onOptionsItemSelected(MenuItem item) {
        // Handle action bar item clicks here. The action bar will
        // automatically handle clicks on the Home/Up button, so long
        // as you specify a parent activity in AndroidManifest.xml.
        int id = item.getItemId();

        //noinspection SimplifiableIfStatement
        if (id == R.id.action_settings) {
            return true;
        }

        return super.onOptionsItemSelected(item);
    }
    public void goCheckin(View view) {
        startActivity(new Intent(this, CheckinActivity.class));
    }

    @Override
    public boolean onQueryTextSubmit(String query) {
        return false;
    }

    @Override
    public boolean onQueryTextChange(String newText) {
        return false;
    }
}
