package ninckblokje.poc.datasets.service;

import io.micronaut.context.event.StartupEvent;
import io.micronaut.runtime.event.annotation.EventListener;
import ninckblokje.poc.datasets.model.Data;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import javax.inject.Singleton;
import java.time.OffsetDateTime;
import java.util.ArrayList;
import java.util.Collections;
import java.util.List;
import java.util.stream.IntStream;

@Singleton
public class DatasetService {

    private static final Logger LOG = LoggerFactory.getLogger(DatasetService.class);

    private List<Data> dataset = Collections.synchronizedList(new ArrayList<>());

    public List<Data> getDataset(int size) {
        LOG.info("Returning dataset copy of size {}", size);
        return dataset.subList(0, size);
    }

    @EventListener
    public void initDataset(StartupEvent startupEvent) {
        LOG.info("Creating dataset");

        IntStream.range(0, 10000000).asLongStream()
                .mapToObj(i -> createData(i))
                .forEach(dataset::add);

        LOG.info("Done creating dataset, size {}", dataset.size());
    }

    Data createData(long id) {
        var data = new Data();

        data.setId(id);
        data.setDateTime(OffsetDateTime.parse("2015-01-01T00:00:00+01:00").plusMinutes(id * 15));

        data.setValue1((id * 5) + 100);
        data.setValue2((id * 2) + 25);
        data.setValue3((id * 3) + 0);

        return data;
    }
}
