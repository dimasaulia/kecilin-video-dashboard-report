async function dataLoader({ url, func }) {
    const resp = await fetcher(`${url}`);
    func(resp);
}
