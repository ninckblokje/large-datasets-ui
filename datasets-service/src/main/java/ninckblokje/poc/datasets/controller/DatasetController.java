package ninckblokje.poc.datasets.controller;

import io.micronaut.http.annotation.Controller;
import io.micronaut.http.annotation.Get;
import io.micronaut.http.annotation.Produces;
import io.micronaut.http.annotation.QueryValue;
import ninckblokje.poc.datasets.model.Data;
import ninckblokje.poc.datasets.service.DatasetService;

import java.util.List;

import static io.micronaut.http.MediaType.APPLICATION_JSON;

@Controller("/api/dataset")
public class DatasetController {

    private final DatasetService datasetService;

    public DatasetController(DatasetService datasetService) {
        this.datasetService = datasetService;
    }

    @Get
    @Produces(APPLICATION_JSON)
    public List<Data> getDataset(@QueryValue(defaultValue = "10") int size) {
        return datasetService.getDataset(size);
    }
}
